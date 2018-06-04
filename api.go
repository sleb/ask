package ask

// PlainText indicates that the response is plain text (not SSML)
const PlainText = "PlainText"

// APIVersion is defined to be "1.0"
const APIVersion = "1.0"

// Session encapsulates attributes of the current conversation with the user.
//
// Standard request types (CanFulfillIntentReqeuest, LaunchRequest, IntentRequest,
// and SessionEndedRequest) include the session object. The GameEngine interface
// includes a session object also.
//
// Requests from interfaces such as AudioPlayer and PlaybackController are not
// sent in the context of a session, so they do not include the session object.
// The context.System.user and context.System.application objects provide the
// same user and application information as the same objects within session
// – see Context Object.
type Session struct {
	New bool `json:"new"`
}

// RequestEnvelope encapsulates the top-level request request object.
// All requests include the version, context, and request objects at the top level.
// The session object is included for all standard requests, but it is not included
// for AudioPlayer, VideoApp, or PlaybackController requests.
type RequestEnvelope struct {
	Version string       `json:"version"`
	Session *Session     `json:"session"`
	Request *interface{} `json:"request"`
	Context *interface{} `json:"context"`
}

// OutputSpeech encapsulates the speech response to the user.
//
// This object is used for setting both the outputSpeech and the reprompt properties.
//
// This object can only be included when sending a response to a CanFulfillIntentRequest,
// LaunchRequest, IntentRequest, or InputHandlerEvent.
type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Response defines what to render to the user and whether to end the current session.
type Response struct {
	ShouldEndSession bool          `json:"shouldEndSession"`
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
}

// ResponseEnvelope encapsulates the top-level response object from the skill.
//
// Note the following size limitations for the response:
//
// The outputSpeech response cannot exceed 8000 characters.
// All of the text included in a card cannot exceed 8000 characters. This includes the title, content, text, and image URLs.
// An image URL (smallImageUrl or largeImageUrl) cannot exceed 2000 characters.
// The token included in an audioItem.stream for the AudioPlayer.Play directive cannot exceed 1024 characters.
// The url included in an audioItem.stream for the AudioPlayer.Play directive cannot exceed 8000 characters.
// The total size of your response cannot exceed 24 kilobytes.
// If your response exceeds these limits, the Alexa service returns an error.
type ResponseEnvelope struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          *Response              `json:"response"`
}
