package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nyaruka/goflow/excellent"
	"github.com/nyaruka/goflow/flows"
	"github.com/nyaruka/goflow/flows/definition"
	"github.com/nyaruka/goflow/flows/engine"
	"github.com/nyaruka/goflow/flows/events"
	"github.com/nyaruka/goflow/flows/runs"
	"github.com/nyaruka/goflow/utils"
)

func main() {
	http.HandleFunc("/flow/start", handleStart)
	http.HandleFunc("/flow/resume", handleResume)
	http.HandleFunc("/migrate", handleMigrate)
	http.HandleFunc("/expression", handleExpression)
	fmt.Println()
	fmt.Println("Server running on port 8080")
	fmt.Println("Endpoints:")
	fmt.Println("  /flow/start    - start a flow. requires a list of flows (first will be started) and contact")
	fmt.Println("  /flow/resume   - resume a flow. requires a contact, list of flows, list of runs and event causing the resumption")
	fmt.Println("  /migrate       - migrates a list of flows")
	fmt.Println("  /expression    - lets you test an expression, requires a context and expression")
	fmt.Println()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//-----------------------------------------------------------------------------
// Execute handler
//-----------------------------------------------------------------------------

type migrateRequest struct {
	Flows json.RawMessage `json:"flows"`
}

func handleMigrate(w http.ResponseWriter, r *http.Request) {
	migrate := migrateRequest{}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, err)
		return
	}

	if err := r.Body.Close(); err != nil {
		writeError(w, err)
		return
	}

	if err := json.Unmarshal(body, &migrate); err != nil {
		writeError(w, err)
		return
	}

	if migrate.Flows == nil {
		writeError(w, fmt.Errorf("Missing flows element"))
		return
	}

	flows, err := definition.ReadLegacyFlows(migrate.Flows)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flows)

}

type flowResponse struct {
	Contact   *flows.Contact  `json:"contact"`
	RunOutput flows.RunOutput `json:"run_output"`
}

type startRequest struct {
	Flows   json.RawMessage `json:"flows"`
	Contact json.RawMessage `json:"contact"`
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	start := startRequest{}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, err)
		return
	}
	if err := r.Body.Close(); err != nil {
		writeError(w, err)
		return
	}
	if err := json.Unmarshal(body, &start); err != nil {
		writeError(w, err)
		return
	}

	if start.Flows == nil || start.Contact == nil {
		writeError(w, fmt.Errorf("Missing contact or flows element"))
		return
	}

	// read our flows
	startFlows, err := definition.ReadFlows(start.Flows)
	if err != nil {
		writeError(w, fmt.Errorf("Error parsing flows: %s", err))
		return
	}

	// read our contact
	contact, err := flows.ReadContact(start.Contact)
	if err != nil {
		writeError(w, fmt.Errorf("Error parsing contact: %s", err))
		return
	}

	// build our environment
	env := engine.NewFlowEnvironment(utils.NewDefaultEnvironment(), startFlows, []flows.FlowRun{}, []*flows.Contact{contact})

	// start our flow
	output, err := engine.StartFlow(env, startFlows[0], contact, nil)
	if err != nil {
		writeError(w, fmt.Errorf("Error starting flow: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := &flowResponse{Contact: contact, RunOutput: output}
	json.NewEncoder(w).Encode(resp)
}

type resumeRequest struct {
	Contact   json.RawMessage      `json:"contact"`
	Flows     json.RawMessage      `json:"flows"`
	RunOutput json.RawMessage      `json:"run_output"`
	Event     *utils.TypedEnvelope `json:"event"`
}

func handleResume(w http.ResponseWriter, r *http.Request) {
	resume := resumeRequest{}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, err)
		return
	}
	if err := r.Body.Close(); err != nil {
		writeError(w, err)
		return
	}
	if err := json.Unmarshal(body, &resume); err != nil {
		writeError(w, err)
		return
	}

	if resume.Flows == nil || resume.RunOutput == nil || resume.Event == nil || resume.Contact == nil {
		writeError(w, fmt.Errorf("Missing flows, run_output, contact or event element"))
		return
	}

	// read our flows
	flowList, err := definition.ReadFlows(resume.Flows)
	if err != nil {
		writeError(w, fmt.Errorf("Error parsing flows: %s", err))
		return
	}

	// read our run
	runOutput, err := runs.ReadRunOutput(resume.RunOutput)
	if err != nil {
		writeError(w, fmt.Errorf("Error parsing run output: %s", err))
		return
	}

	// our contact
	contact, err := flows.ReadContact(resume.Contact)
	if err != nil {
		writeError(w, fmt.Errorf("Error parsing run contact: %s", err))
		return
	}

	// and our event
	event, err := events.EventFromEnvelope(resume.Event)
	if err != nil {
		writeError(w, fmt.Errorf("Error reading event: %s", err))
		return
	}

	// build our environment
	env := engine.NewFlowEnvironment(utils.NewDefaultEnvironment(), flowList, runOutput.Runs(), []*flows.Contact{contact})

	// hydrate all our runs
	for _, run := range runOutput.Runs() {
		run.Hydrate(env)
	}

	// set our contact on our run
	activeRun := runOutput.ActiveRun()

	// resume our flow
	output, err := engine.ResumeFlow(env, activeRun, event)
	if err != nil {
		writeError(w, fmt.Errorf("Error resuming flow: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := &flowResponse{Contact: contact, RunOutput: output}
	json.NewEncoder(w).Encode(resp)
}

//-----------------------------------------------------------------------------
// Expression handler
//-----------------------------------------------------------------------------

type expressionResponse struct {
	Result string   `json:"result"`
	Errors []string `json:"errors"`
}

type expressionRequest struct {
	Expression string          `json:"expression"`
	Context    json.RawMessage `json:"context"`
}

func handleExpression(w http.ResponseWriter, r *http.Request) {
	expression := expressionRequest{}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, err)
		return
	}
	if err := r.Body.Close(); err != nil {
		writeError(w, err)
		return
	}
	if err := json.Unmarshal(body, &expression); err != nil {
		writeError(w, err)
		return
	}

	if expression.Context == nil || expression.Expression == "" {
		writeError(w, fmt.Errorf("Missing context or expression element"))
		return
	}

	// build up our context
	context, err := runs.ReadContext(expression.Context)
	if err != nil {
		writeError(w, err)
		return
	}

	// evaluate it
	result, err := excellent.EvaluateTemplateAsString(utils.NewDefaultEnvironment(), context, expression.Expression)

	w.WriteHeader(http.StatusOK)
	response := expressionResponse{Result: result}
	if err != nil {
		switch err.(type) {
		case excellent.TemplateErrors:
			templateErrs := err.(excellent.TemplateErrors)
			errs := make([]string, len(templateErrs))
			for i := range errs {
				errs[i] = templateErrs[i].Error()
			}
			response.Errors = errs
		default:
			response.Errors = []string{err.Error()}
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

//-----------------------------------------------------------------------------
// Error utils
//-----------------------------------------------------------------------------

type errorResponse struct {
	Text string `json:"error"`
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse{err.Error()})
}
