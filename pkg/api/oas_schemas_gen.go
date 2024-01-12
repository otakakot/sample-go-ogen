// Code generated by ogen, DO NOT EDIT.

package api

// Ref: #/components/schemas/ErrorResponseSchema
type ErrorResponseSchema struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *ErrorResponseSchema) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *ErrorResponseSchema) SetMessage(val string) {
	s.Message = val
}

func (*ErrorResponseSchema) getHealthRes()  {}
func (*ErrorResponseSchema) postHealthRes() {}

// Ref: #/components/schemas/HealthRequestSchema
type HealthRequestSchema struct {
	// Message.
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *HealthRequestSchema) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *HealthRequestSchema) SetMessage(val string) {
	s.Message = val
}

// Ref: #/components/schemas/HealthResponseSchema
type HealthResponseSchema struct {
	// Message.
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *HealthResponseSchema) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *HealthResponseSchema) SetMessage(val string) {
	s.Message = val
}

func (*HealthResponseSchema) getHealthRes()  {}
func (*HealthResponseSchema) postHealthRes() {}
