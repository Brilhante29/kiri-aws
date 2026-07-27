package gamelift

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/Brilhante29/kiri-aws/internal/server"
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// handlerFunc is a type alias for handler functions.
type handlerFunc func(http.ResponseWriter, *http.Request)

// getActionHandlers returns a map of action names to handler functions.
func (s *Service) getActionHandlers() map[string]handlerFunc {
	return map[string]handlerFunc{
		// Build operations
		"CreateBuild":   s.CreateBuild,
		"DescribeBuild": s.DescribeBuild,
		"ListBuilds":    s.ListBuilds,
		"DeleteBuild":   s.DeleteBuild,
		// Fleet operations
		"CreateFleet":             s.CreateFleet,
		"DescribeFleetAttributes": s.DescribeFleetAttributes,
		"ListFleets":              s.ListFleets,
		"DeleteFleet":             s.DeleteFleet,
		// Game session operations
		"CreateGameSession":    s.CreateGameSession,
		"DescribeGameSessions": s.DescribeGameSessions,
		"UpdateGameSession":    s.UpdateGameSession,
		// Player session operations
		"CreatePlayerSession":    s.CreatePlayerSession,
		"CreatePlayerSessions":   s.CreatePlayerSessions,
		"DescribePlayerSessions": s.DescribePlayerSessions,
	}
}

// DispatchAction dispatches the request to the appropriate handler.
func (s *Service) DispatchAction(w http.ResponseWriter, r *http.Request) {
	target := r.Header.Get("X-Amz-Target")
	action := strings.TrimPrefix(target, "GameLift.")

	handlers := s.getActionHandlers()
	if handler, ok := handlers[action]; ok {
		handler(w, r)

		return
	}

	writeError(w, r,"UnknownOperationException", "The operation "+action+" is not valid.", http.StatusBadRequest)
}

// CreateBuild handles the CreateBuild API.
func (s *Service) CreateBuild(w http.ResponseWriter, r *http.Request) {
	var req CreateBuildRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	build, err := s.storage.CreateBuild(r.Context(), &req)
	if err != nil {
		handleError(w, r, err)

		return
	}

	resp := &CreateBuildResponse{
		Build: convertToBuildOutput(build),
	}

	writeResponse(w, r, resp)
}

// DescribeBuild handles the DescribeBuild API.
func (s *Service) DescribeBuild(w http.ResponseWriter, r *http.Request) {
	var req DescribeBuildRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	if req.BuildID == "" {
		writeError(w, r,"InvalidRequestException", "BuildId is required", http.StatusBadRequest)

		return
	}

	build, err := s.storage.DescribeBuild(r.Context(), req.BuildID)
	if err != nil {
		handleError(w, r, err)

		return
	}

	resp := &DescribeBuildResponse{
		Build: convertToBuildOutput(build),
	}

	writeResponse(w, r, resp)
}

// ListBuilds handles the ListBuilds API.
func (s *Service) ListBuilds(w http.ResponseWriter, r *http.Request) {
	var req ListBuildsRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	limit := int32(100)
	if req.Limit != nil && *req.Limit > 0 {
		limit = *req.Limit
	}

	builds, err := s.storage.ListBuilds(r.Context(), req.Status, limit)
	if err != nil {
		handleError(w, r, err)

		return
	}

	buildOutputs := make([]BuildOutput, 0, len(builds))
	for _, build := range builds {
		buildOutputs = append(buildOutputs, *convertToBuildOutput(build))
	}

	resp := &ListBuildsResponse{
		Builds: buildOutputs,
	}

	writeResponse(w, r, resp)
}

// DeleteBuild handles the DeleteBuild API.
func (s *Service) DeleteBuild(w http.ResponseWriter, r *http.Request) {
	var req DeleteBuildRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	if req.BuildID == "" {
		writeError(w, r,"InvalidRequestException", "BuildId is required", http.StatusBadRequest)

		return
	}

	if err := s.storage.DeleteBuild(r.Context(), req.BuildID); err != nil {
		handleError(w, r, err)

		return
	}

	writeResponse(w, r, &DeleteBuildResponse{})
}

// CreateFleet handles the CreateFleet API.
func (s *Service) CreateFleet(w http.ResponseWriter, r *http.Request) {
	var req CreateFleetRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	fleet, err := s.storage.CreateFleet(r.Context(), &req)
	if err != nil {
		handleError(w, r, err)

		return
	}

	resp := &CreateFleetResponse{
		FleetAttributes: convertToFleetAttributesOutput(fleet),
	}

	writeResponse(w, r, resp)
}

// DescribeFleetAttributes handles the DescribeFleetAttributes API.
func (s *Service) DescribeFleetAttributes(w http.ResponseWriter, r *http.Request) {
	var req DescribeFleetAttributesRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	fleets, err := s.storage.DescribeFleetAttributes(r.Context(), req.FleetIDs)
	if err != nil {
		handleError(w, r, err)

		return
	}

	fleetOutputs := make([]FleetAttributesOutput, 0, len(fleets))
	for _, fleet := range fleets {
		fleetOutputs = append(fleetOutputs, *convertToFleetAttributesOutput(fleet))
	}

	resp := &DescribeFleetAttributesResponse{
		FleetAttributes: fleetOutputs,
	}

	writeResponse(w, r, resp)
}

// ListFleets handles the ListFleets API.
func (s *Service) ListFleets(w http.ResponseWriter, r *http.Request) {
	var req ListFleetsRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	limit := int32(100)
	if req.Limit != nil && *req.Limit > 0 {
		limit = *req.Limit
	}

	fleetIDs, err := s.storage.ListFleets(r.Context(), req.BuildID, limit)
	if err != nil {
		handleError(w, r, err)

		return
	}

	resp := &ListFleetsResponse{
		FleetIDs: fleetIDs,
	}

	writeResponse(w, r, resp)
}

// DeleteFleet handles the DeleteFleet API.
func (s *Service) DeleteFleet(w http.ResponseWriter, r *http.Request) {
	var req DeleteFleetRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	if req.FleetID == "" {
		writeError(w, r,"InvalidRequestException", "FleetId is required", http.StatusBadRequest)

		return
	}

	if err := s.storage.DeleteFleet(r.Context(), req.FleetID); err != nil {
		handleError(w, r, err)

		return
	}

	writeResponse(w, r, &DeleteFleetResponse{})
}

// CreateGameSession handles the CreateGameSession API.
func (s *Service) CreateGameSession(w http.ResponseWriter, r *http.Request) {
	var req CreateGameSessionRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	gameSession, err := s.storage.CreateGameSession(r.Context(), &req)
	if err != nil {
		handleError(w, r, err)

		return
	}

	resp := &CreateGameSessionResponse{
		GameSession: convertToGameSessionOutput(gameSession),
	}

	writeResponse(w, r, resp)
}

// DescribeGameSessions handles the DescribeGameSessions API.
func (s *Service) DescribeGameSessions(w http.ResponseWriter, r *http.Request) {
	var req DescribeGameSessionsRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	sessions, err := s.storage.DescribeGameSessions(r.Context(), req.FleetID, req.GameSessionID)
	if err != nil {
		handleError(w, r, err)

		return
	}

	sessionOutputs := make([]GameSessionOutput, 0, len(sessions))
	for _, session := range sessions {
		sessionOutputs = append(sessionOutputs, *convertToGameSessionOutput(session))
	}

	resp := &DescribeGameSessionsResponse{
		GameSessions: sessionOutputs,
	}

	writeResponse(w, r, resp)
}

// UpdateGameSession handles the UpdateGameSession API.
func (s *Service) UpdateGameSession(w http.ResponseWriter, r *http.Request) {
	var req UpdateGameSessionRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	if req.GameSessionID == "" {
		writeError(w, r,"InvalidRequestException", "GameSessionId is required", http.StatusBadRequest)

		return
	}

	gameSession, err := s.storage.UpdateGameSession(r.Context(), &req)
	if err != nil {
		handleError(w, r, err)

		return
	}

	resp := &UpdateGameSessionResponse{
		GameSession: convertToGameSessionOutput(gameSession),
	}

	writeResponse(w, r, resp)
}

// CreatePlayerSession handles the CreatePlayerSession API.
func (s *Service) CreatePlayerSession(w http.ResponseWriter, r *http.Request) {
	var req CreatePlayerSessionRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	if req.GameSessionID == "" {
		writeError(w, r,"InvalidRequestException", "GameSessionId is required", http.StatusBadRequest)

		return
	}

	if req.PlayerID == "" {
		writeError(w, r,"InvalidRequestException", "PlayerId is required", http.StatusBadRequest)

		return
	}

	playerSession, err := s.storage.CreatePlayerSession(r.Context(), req.GameSessionID, req.PlayerID)
	if err != nil {
		handleError(w, r, err)

		return
	}

	resp := &CreatePlayerSessionResponse{
		PlayerSession: convertToPlayerSessionOutput(playerSession),
	}

	writeResponse(w, r, resp)
}

// CreatePlayerSessions handles the CreatePlayerSessions API.
func (s *Service) CreatePlayerSessions(w http.ResponseWriter, r *http.Request) {
	var req CreatePlayerSessionsRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	if req.GameSessionID == "" {
		writeError(w, r,"InvalidRequestException", "GameSessionId is required", http.StatusBadRequest)

		return
	}

	if len(req.PlayerIDs) == 0 {
		writeError(w, r,"InvalidRequestException", "PlayerIds is required", http.StatusBadRequest)

		return
	}

	playerSessions, err := s.storage.CreatePlayerSessions(r.Context(), req.GameSessionID, req.PlayerIDs)
	if err != nil {
		handleError(w, r, err)

		return
	}

	sessionOutputs := make([]PlayerSessionOutput, 0, len(playerSessions))
	for _, session := range playerSessions {
		sessionOutputs = append(sessionOutputs, *convertToPlayerSessionOutput(session))
	}

	resp := &CreatePlayerSessionsResponse{
		PlayerSessions: sessionOutputs,
	}

	writeResponse(w, r, resp)
}

// DescribePlayerSessions handles the DescribePlayerSessions API.
func (s *Service) DescribePlayerSessions(w http.ResponseWriter, r *http.Request) {
	var req DescribePlayerSessionsRequest
	if err := decodeRequest(r, &req); err != nil {
		writeError(w, r,"InvalidRequestException", "Invalid request body", http.StatusBadRequest)

		return
	}

	sessions, err := s.storage.DescribePlayerSessions(r.Context(), req.GameSessionID, req.PlayerSessionID, req.PlayerID)
	if err != nil {
		handleError(w, r, err)

		return
	}

	sessionOutputs := make([]PlayerSessionOutput, 0, len(sessions))
	for _, session := range sessions {
		sessionOutputs = append(sessionOutputs, *convertToPlayerSessionOutput(session))
	}

	resp := &DescribePlayerSessionsResponse{
		PlayerSessions: sessionOutputs,
	}

	writeResponse(w, r, resp)
}

// Helper functions.

// convertToBuildOutput converts a Build to BuildOutput.
func convertToBuildOutput(build *Build) *BuildOutput {
	return &BuildOutput{
		BuildID:         build.BuildID,
		BuildARN:        build.BuildARN,
		Name:            build.Name,
		Version:         build.Version,
		Status:          build.Status,
		SizeOnDisk:      build.SizeOnDisk,
		OperatingSystem: build.OperatingSystem,
		CreationTime:    timePtr(build.CreationTime),
	}
}

// convertToFleetAttributesOutput converts a Fleet to FleetAttributesOutput.
func convertToFleetAttributesOutput(fleet *Fleet) *FleetAttributesOutput {
	return &FleetAttributesOutput{
		FleetID:                        fleet.FleetID,
		FleetARN:                       fleet.FleetARN,
		FleetType:                      fleet.FleetType,
		InstanceType:                   fleet.InstanceType,
		Description:                    fleet.Description,
		Name:                           fleet.Name,
		CreationTime:                   timePtr(fleet.CreationTime),
		Status:                         fleet.Status,
		BuildID:                        fleet.BuildID,
		ServerLaunchPath:               fleet.ServerLaunchPath,
		NewGameSessionProtectionPolicy: fleet.NewGameSessionProtectionPolicy,
	}
}

// convertToGameSessionOutput converts a GameSession to GameSessionOutput.
func convertToGameSessionOutput(session *GameSession) *GameSessionOutput {
	return &GameSessionOutput{
		GameSessionID:             session.GameSessionID,
		Name:                      session.Name,
		FleetID:                   session.FleetID,
		FleetARN:                  session.FleetARN,
		CreationTime:              timePtr(session.CreationTime),
		CurrentPlayerSessionCount: session.CurrentPlayerSessionCount,
		MaximumPlayerSessionCount: session.MaximumPlayerSessionCount,
		Status:                    session.Status,
		IPAddress:                 session.IPAddress,
		Port:                      session.Port,
	}
}

// convertToPlayerSessionOutput converts a PlayerSession to PlayerSessionOutput.
func convertToPlayerSessionOutput(session *PlayerSession) *PlayerSessionOutput {
	return &PlayerSessionOutput{
		PlayerSessionID: session.PlayerSessionID,
		PlayerID:        session.PlayerID,
		GameSessionID:   session.GameSessionID,
		FleetID:         session.FleetID,
		FleetARN:        session.FleetARN,
		CreationTime:    timePtr(session.CreationTime),
		Status:          session.Status,
		IPAddress:       session.IPAddress,
		Port:            session.Port,
	}
}

func timePtr(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}

	return &t
}

// isCBOR checks if the request uses Smithy RPC v2 CBOR protocol.
func isCBOR(r *http.Request) bool {
	return r.Header.Get("smithy-protocol") == "rpc-v2-cbor" || r.Header.Get("Content-Type") == "application/cbor"
}

// decodeRequest decodes a request body using the appropriate codec.
func decodeRequest(r *http.Request, v any) error {
	if isCBOR(r) {
		return server.DecodeCBORRequest(r, v)
	}

	return json.NewDecoder(r.Body).Decode(v)
}

// writeResponse writes a response using the appropriate codec.
func writeResponse(w http.ResponseWriter, r *http.Request, resp any) {
	if isCBOR(r) {
		server.WriteCBORResponse(w, resp)

		return
	}

	w.Header().Set("Content-Type", "application/x-amz-json-1.0")
	w.Header().Set("x-amzn-RequestId", uuid.New().String())
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// writeError writes an error response using the appropriate codec.
func writeError(w http.ResponseWriter, r *http.Request, code, message string, status int) {
	if isCBOR(r) {
		server.WriteCBORError(w, code, message, status)

		return
	}

	service.WriteJSONError(w, service.ContentTypeAmzJSON10, code, message, status)
}

// handleError handles service errors using the appropriate codec.
func handleError(w http.ResponseWriter, r *http.Request, err error) {
	var glErr *Error
	if errors.As(err, &glErr) {
		writeError(w, r, glErr.Code, glErr.Message, getErrorStatus(glErr.Code))

		return
	}

	writeError(w, r, "InternalServiceException", err.Error(), http.StatusInternalServerError)
}

// getErrorStatus returns the HTTP status code for a given error code.
func getErrorStatus(code string) int {
	switch code {
	case errNotFoundException:
		return http.StatusNotFound
	case errInvalidRequestException:
		return http.StatusBadRequest
	case errConflictException:
		return http.StatusConflict
	case errLimitExceededException:
		return http.StatusTooManyRequests
	default:
		return http.StatusBadRequest
	}
}
