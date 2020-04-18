// Copyright 2020 Justice Nanhou. All rights reserved.
// license that can be found in the LICENSE file.

package message

// http://www.fipa.org/specs/fipa00061/SC00061G.html

// Message represent a message object send between agents
type Message struct {
	Type           string // Performative :Type of communicative acts. http://www.fipa.org/specs/fipa00037/SC00037J.html
	Sender         string // Participant in communication
	Receiver       string // Participant in communication
	replyTo        string // Participant in communication
	Content        string // Content of message
	language       string // Description of Content
	encoding       string // Description of Content
	ontology       string // Description of Content
	protocol       string // Control of conversation
	conversationID string // Control of conversation
	replyWith      string // Control of conversation
	inReplyTo      string // Control of conversation
	replyBy        string // Control of conversation
}

// They are listed in the FIPA Communicative Act Library Specification.
const (
	// TypeAcceptProposal The action of accepting a previously submitted propose to perform an action.
	TypeAcceptProposal string = "accept-proposal"
	// The action of agreeing to perform a requestd action made by another agent. Agent will carry it out.
	TypeAgree string = "agree"
	// Agent wants to cancel a previous request.
	TypeCancel string = "cancel"
	// Agent issues a call for proposals. It contains the actions to be carried out and any other terms of the agreement.
	TypeCfp string = "cfp"
	// The sender confirms to the receiver the truth of the content. The sender initially believed that the receiver was unsure about it.
	TypeConfirm string = "confirm"
	// The sender confirms to the receiver the falsity of the content.
	TypeDisconfirm string = "disconfirm"
	// Tell the other agent that a previously requested action failed.
	TypeFailure string = "failure"
	// Tell another agent something. The sender must believe in the truth of the statement. Most used performative.
	TypeInform string = "inform"
	// Used as content of request to ask another agent to tell us is a statement is true or false.
	TypeIinformIf string = "inform-if"
	// Like inform-if but asks for the value of the expression.
	TypeInformRef string = "inform-ref"
	// Sent when the agent did not understand the message.
	TypeNotUnderstood string = "not-understood"
	// Asks another agent so forward this same propagate message to others.
	TypePropagate string = "propagate"
	// Used as a response to a cfp. Agent proposes a deal.
	TypePropose string = "propose"
	// The sender wants the receiver to select target agents denoted by a given description and to send an embedded message to them.
	TypeProxy string = "proxy"
	// The action of asking another agent whether or not a given proposition is true.
	TypeQueryIf string = "query-if"
	// The action of asking another agent for the object referred to by an referential expression.
	TypeQueryRef string = "query-ref"
	// The action of refusing to perform a given action, and explaining the reason for the refusal.
	TypeRefuse string = "refuse"
	// The action of rejecting a proposal to perform some action during a negotiation.
	TypeRejectProposal string = "reject-proposal"
	// The sender requests the receiver to perform some action. Usually to request the receiver to perform another communicative act.
	TypeRequest string = "request"
	// The sender wants the receiver to perform some action when some given proposition becomes true.
	TypeRequestWhen string = "request-when"
	// The sender wants the receiver to perform some action as soon as some proposition becomes true and thereafter each time the proposition becomes true again.
	TypeRequestWhenever string = "request-whenever"
	// The act of requesting a persistent intention to notify the sender of the value of a reference, and to notify again whenever the object identified by the reference changes.
	TypeSubscribe string = "subscribe"
)
