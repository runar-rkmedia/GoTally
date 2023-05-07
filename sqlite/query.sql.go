// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package sqlite

import (
	"context"
	"database/sql"
	"time"
)

const getAllGames = `-- name: GetAllGames :many
SELECT id, created_at, updated_at, name, description, user_id, rule_id, based_on_game, template_id, score, moves, play_state, data, data_at_start, history from game
`

func (q *Queries) GetAllGames(ctx context.Context) ([]Game, error) {
	rows, err := q.db.QueryContext(ctx, getAllGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Game
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.UserID,
			&i.RuleID,
			&i.BasedOnGame,
			&i.TemplateID,
			&i.Score,
			&i.Moves,
			&i.PlayState,
			&i.Data,
			&i.DataAtStart,
			&i.History,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllRules = `-- name: GetAllRules :many
SELECT id, slug, created_at, updated_at, mode, description, size_x, size_y, max_moves, target_cell_value, target_score, recreate_on_swipe, no_reswipe, no_multiply, no_addition from rule
`

func (q *Queries) GetAllRules(ctx context.Context) ([]Rule, error) {
	rows, err := q.db.QueryContext(ctx, getAllRules)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rule
	for rows.Next() {
		var i Rule
		if err := rows.Scan(
			&i.ID,
			&i.Slug,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Mode,
			&i.Description,
			&i.SizeX,
			&i.SizeY,
			&i.MaxMoves,
			&i.TargetCellValue,
			&i.TargetScore,
			&i.RecreateOnSwipe,
			&i.NoReswipe,
			&i.NoMultiply,
			&i.NoAddition,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSessions = `-- name: GetAllSessions :many
SELECT id, created_at, updated_at, invalid_after, user_id from session
`

func (q *Queries) GetAllSessions(ctx context.Context) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, getAllSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.InvalidAfter,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTemplates = `-- name: GetAllTemplates :many
SELECT id, created_at, updated_at, rule_id, created_by, updated_by, name, description, challenge_number, ideal_moves, ideal_score, data from game_template
`

func (q *Queries) GetAllTemplates(ctx context.Context) ([]GameTemplate, error) {
	rows, err := q.db.QueryContext(ctx, getAllTemplates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GameTemplate
	for rows.Next() {
		var i GameTemplate
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RuleID,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.Name,
			&i.Description,
			&i.ChallengeNumber,
			&i.IdealMoves,
			&i.IdealScore,
			&i.Data,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, created_at, updated_at, username, active_game_id from user
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Username,
			&i.ActiveGameID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChallengeStatsForUser = `-- name: GetChallengeStatsForUser :many
SELECT
	g.id as game_id
	, g.template_id
	, g.score
	, g.moves
	, u.id as user_id
	, u.username 
FROM
	game AS g
	JOIN user AS u ON u.id = g.user_id 
WHERE
	template_id IS NOT NULL
  AND user_id = ?
`

type GetChallengeStatsForUserRow struct {
	GameID     string
	TemplateID sql.NullString
	Score      int64
	Moves      int64
	UserID     string
	Username   string
}

func (q *Queries) GetChallengeStatsForUser(ctx context.Context, userID string) ([]GetChallengeStatsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getChallengeStatsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetChallengeStatsForUserRow
	for rows.Next() {
		var i GetChallengeStatsForUserRow
		if err := rows.Scan(
			&i.GameID,
			&i.TemplateID,
			&i.Score,
			&i.Moves,
			&i.UserID,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGame = `-- name: GetGame :one
select id, created_at, updated_at, name, description, user_id, rule_id, based_on_game, template_id, score, moves, play_state, data, data_at_start, history from game
where id == ?
`

func (q *Queries) GetGame(ctx context.Context, id string) (Game, error) {
	row := q.db.QueryRowContext(ctx, getGame, id)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.UserID,
		&i.RuleID,
		&i.BasedOnGame,
		&i.TemplateID,
		&i.Score,
		&i.Moves,
		&i.PlayState,
		&i.Data,
		&i.DataAtStart,
		&i.History,
	)
	return i, err
}

const getGameChallengesTemplates = `-- name: GetGameChallengesTemplates :many
select id, created_at, updated_at, rule_id, created_by, updated_by, name, description, challenge_number, ideal_moves, ideal_score, data from game_template
where challenge_number is not null
order by challenge_number
`

func (q *Queries) GetGameChallengesTemplates(ctx context.Context) ([]GameTemplate, error) {
	rows, err := q.db.QueryContext(ctx, getGameChallengesTemplates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GameTemplate
	for rows.Next() {
		var i GameTemplate
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RuleID,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.Name,
			&i.Description,
			&i.ChallengeNumber,
			&i.IdealMoves,
			&i.IdealScore,
			&i.Data,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGameTemplate = `-- name: GetGameTemplate :one
select id, created_at, updated_at, rule_id, created_by, updated_by, name, description, challenge_number, ideal_moves, ideal_score, data from game_template
where id = ?
`

func (q *Queries) GetGameTemplate(ctx context.Context, id string) (GameTemplate, error) {
	row := q.db.QueryRowContext(ctx, getGameTemplate, id)
	var i GameTemplate
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RuleID,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.Name,
		&i.Description,
		&i.ChallengeNumber,
		&i.IdealMoves,
		&i.IdealScore,
		&i.Data,
	)
	return i, err
}

const getGameTemplateByChallengeNumber = `-- name: GetGameTemplateByChallengeNumber :one
select id, created_at, updated_at, rule_id, created_by, updated_by, name, description, challenge_number, ideal_moves, ideal_score, data from game_template
where challenge_number = ?
`

func (q *Queries) GetGameTemplateByChallengeNumber(ctx context.Context, challengeNumber sql.NullInt64) (GameTemplate, error) {
	row := q.db.QueryRowContext(ctx, getGameTemplateByChallengeNumber, challengeNumber)
	var i GameTemplate
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RuleID,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.Name,
		&i.Description,
		&i.ChallengeNumber,
		&i.IdealMoves,
		&i.IdealScore,
		&i.Data,
	)
	return i, err
}

const getOriginalGame = `-- name: GetOriginalGame :one
SELECT o.id, o.created_at, o.updated_at, o.name, o.description, o.user_id, o.rule_id, o.based_on_game, o.template_id, o.score, o.moves, o.play_state, o.data, o.data_at_start, o.history
  FROM game AS g 
    JOIN game o ON o.id == g.based_on_game
 WHERE g.id = '?'
`

func (q *Queries) GetOriginalGame(ctx context.Context) (Game, error) {
	row := q.db.QueryRowContext(ctx, getOriginalGame)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.UserID,
		&i.RuleID,
		&i.BasedOnGame,
		&i.TemplateID,
		&i.Score,
		&i.Moves,
		&i.PlayState,
		&i.Data,
		&i.DataAtStart,
		&i.History,
	)
	return i, err
}

const getRule = `-- name: GetRule :one
;
select id, slug, created_at, updated_at, mode, description, size_x, size_y, max_moves, target_cell_value, target_score, recreate_on_swipe, no_reswipe, no_multiply, no_addition from rule
where id == ? or slug == ?
`

type GetRuleParams struct {
	ID   string
	Slug string
}

func (q *Queries) GetRule(ctx context.Context, arg GetRuleParams) (Rule, error) {
	row := q.db.QueryRowContext(ctx, getRule, arg.ID, arg.Slug)
	var i Rule
	err := row.Scan(
		&i.ID,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Mode,
		&i.Description,
		&i.SizeX,
		&i.SizeY,
		&i.MaxMoves,
		&i.TargetCellValue,
		&i.TargetScore,
		&i.RecreateOnSwipe,
		&i.NoReswipe,
		&i.NoMultiply,
		&i.NoAddition,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
select id, created_at, updated_at, username, active_game_id from user
where id == ?
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.ActiveGameID,
	)
	return i, err
}

const getUserBySessionID = `-- name: GetUserBySessionID :one
SELECT
       session.id session_id,
       session.created_at session_created_at,
       session.invalid_after session_invalid_after,

       user.id user_id,
       user.created_at user_created_at,
       user.updated_at user_updated_at,
       user.username,

       game.id game_id,
       game.created_at game_created_at,
       game.updated_at game_updated_at,
       game.description game_description,
       game.name game_Name,
       game.data game_data,
       game.history game_history,
       game.play_state game_play_state,
       game.score game_score,
       game.moves game_moves,
       game.rule_id rule_id

FROM session session
         INNER JOIN user user on user.id = session.user_id
         INNER JOIN game game on user.active_game_id = game.id
WHERE session.id = ? LIMIT 1
`

type GetUserBySessionIDRow struct {
	ID           string
	CreatedAt    time.Time
	InvalidAfter time.Time
	ID_2         string
	CreatedAt_2  time.Time
	UpdatedAt    sql.NullTime
	Username     string
	ID_3         string
	CreatedAt_3  time.Time
	UpdatedAt_2  sql.NullTime
	Description  sql.NullString
	Name         sql.NullString
	Data         []byte
	History      []byte
	PlayState    int64
	Score        int64
	Moves        int64
	RuleID       string
}

func (q *Queries) GetUserBySessionID(ctx context.Context, id string) (GetUserBySessionIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserBySessionID, id)
	var i GetUserBySessionIDRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.InvalidAfter,
		&i.ID_2,
		&i.CreatedAt_2,
		&i.UpdatedAt,
		&i.Username,
		&i.ID_3,
		&i.CreatedAt_3,
		&i.UpdatedAt_2,
		&i.Description,
		&i.Name,
		&i.Data,
		&i.History,
		&i.PlayState,
		&i.Score,
		&i.Moves,
		&i.RuleID,
	)
	return i, err
}

const inserTemplate = `-- name: InserTemplate :one
INSERT INTO game_template
(id, created_at, rule_id, created_by, name, description, challenge_number, ideal_moves, ideal_score, data)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, created_at, updated_at, rule_id, created_by, updated_by, name, description, challenge_number, ideal_moves, ideal_score, data
`

type InserTemplateParams struct {
	ID              string
	CreatedAt       time.Time
	RuleID          string
	CreatedBy       string
	Name            string
	Description     sql.NullString
	ChallengeNumber sql.NullInt64
	IdealMoves      sql.NullInt64
	IdealScore      sql.NullInt64
	Data            []byte
}

func (q *Queries) InserTemplate(ctx context.Context, arg InserTemplateParams) (GameTemplate, error) {
	row := q.db.QueryRowContext(ctx, inserTemplate,
		arg.ID,
		arg.CreatedAt,
		arg.RuleID,
		arg.CreatedBy,
		arg.Name,
		arg.Description,
		arg.ChallengeNumber,
		arg.IdealMoves,
		arg.IdealScore,
		arg.Data,
	)
	var i GameTemplate
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RuleID,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.Name,
		&i.Description,
		&i.ChallengeNumber,
		&i.IdealMoves,
		&i.IdealScore,
		&i.Data,
	)
	return i, err
}

const insertGame = `-- name: InsertGame :one
INSERT INTO game
(id, created_at, updated_at, name, description, user_id, rule_id, score, moves, play_state, data, data_at_start, history, template_id, based_on_game)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, created_at, updated_at, name, description, user_id, rule_id, based_on_game, template_id, score, moves, play_state, data, data_at_start, history
`

type InsertGameParams struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
	Name        sql.NullString
	Description sql.NullString
	UserID      string
	RuleID      string
	Score       int64
	Moves       int64
	PlayState   int64
	Data        []byte
	DataAtStart []byte
	History     []byte
	TemplateID  sql.NullString
	BasedOnGame sql.NullString
}

func (q *Queries) InsertGame(ctx context.Context, arg InsertGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, insertGame,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Description,
		arg.UserID,
		arg.RuleID,
		arg.Score,
		arg.Moves,
		arg.PlayState,
		arg.Data,
		arg.DataAtStart,
		arg.History,
		arg.TemplateID,
		arg.BasedOnGame,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.UserID,
		&i.RuleID,
		&i.BasedOnGame,
		&i.TemplateID,
		&i.Score,
		&i.Moves,
		&i.PlayState,
		&i.Data,
		&i.DataAtStart,
		&i.History,
	)
	return i, err
}

const insertRule = `-- name: InsertRule :one
INSERT INTO rule
(id, slug, created_at, updated_at, description, mode, size_x, size_y, recreate_on_swipe, no_reswipe, no_multiply, no_addition, max_moves, target_cell_value, target_score)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, slug, created_at, updated_at, mode, description, size_x, size_y, max_moves, target_cell_value, target_score, recreate_on_swipe, no_reswipe, no_multiply, no_addition
`

type InsertRuleParams struct {
	ID              string
	Slug            string
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
	Description     sql.NullString
	Mode            int64
	SizeX           int64
	SizeY           int64
	RecreateOnSwipe bool
	NoReswipe       bool
	NoMultiply      bool
	NoAddition      bool
	MaxMoves        sql.NullInt64
	TargetCellValue sql.NullInt64
	TargetScore     sql.NullInt64
}

func (q *Queries) InsertRule(ctx context.Context, arg InsertRuleParams) (Rule, error) {
	row := q.db.QueryRowContext(ctx, insertRule,
		arg.ID,
		arg.Slug,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Description,
		arg.Mode,
		arg.SizeX,
		arg.SizeY,
		arg.RecreateOnSwipe,
		arg.NoReswipe,
		arg.NoMultiply,
		arg.NoAddition,
		arg.MaxMoves,
		arg.TargetCellValue,
		arg.TargetScore,
	)
	var i Rule
	err := row.Scan(
		&i.ID,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Mode,
		&i.Description,
		&i.SizeX,
		&i.SizeY,
		&i.MaxMoves,
		&i.TargetCellValue,
		&i.TargetScore,
		&i.RecreateOnSwipe,
		&i.NoReswipe,
		&i.NoMultiply,
		&i.NoAddition,
	)
	return i, err
}

const insertSession = `-- name: InsertSession :one
INSERT INTO session
    (id, created_at, updated_at, invalid_after, user_id)
VALUES (?, ?, ?, ?, ?)
RETURNING id, created_at, updated_at, invalid_after, user_id
`

type InsertSessionParams struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
	InvalidAfter time.Time
	UserID       string
}

func (q *Queries) InsertSession(ctx context.Context, arg InsertSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, insertSession,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.InvalidAfter,
		arg.UserID,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.InvalidAfter,
		&i.UserID,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO user
    (id, created_at, updated_at, username, active_game_id)
VALUES (?, ?, ?, ?, ?)
RETURNING id, created_at, updated_at, username, active_game_id
`

type InsertUserParams struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
	Username     string
	ActiveGameID string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Username,
		arg.ActiveGameID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.ActiveGameID,
	)
	return i, err
}

const setActiveGameFormUser = `-- name: SetActiveGameFormUser :one
UPDATE user
SET updated_at = ?,
    active_game_id = ?
WHERE id = ?
RETURNING id, created_at, updated_at, username, active_game_id
`

type SetActiveGameFormUserParams struct {
	UpdatedAt    sql.NullTime
	ActiveGameID string
	ID           string
}

func (q *Queries) SetActiveGameFormUser(ctx context.Context, arg SetActiveGameFormUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, setActiveGameFormUser, arg.UpdatedAt, arg.ActiveGameID, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.ActiveGameID,
	)
	return i, err
}

const setPlayStateForGame = `-- name: SetPlayStateForGame :one
UPDATE game
SET updated_at = ?,
    play_state = ?
WHERE id = ?
RETURNING id, created_at, updated_at, name, description, user_id, rule_id, based_on_game, template_id, score, moves, play_state, data, data_at_start, history
`

type SetPlayStateForGameParams struct {
	UpdatedAt sql.NullTime
	PlayState int64
	ID        string
}

func (q *Queries) SetPlayStateForGame(ctx context.Context, arg SetPlayStateForGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, setPlayStateForGame, arg.UpdatedAt, arg.PlayState, arg.ID)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.UserID,
		&i.RuleID,
		&i.BasedOnGame,
		&i.TemplateID,
		&i.Score,
		&i.Moves,
		&i.PlayState,
		&i.Data,
		&i.DataAtStart,
		&i.History,
	)
	return i, err
}

const stats = `-- name: Stats :one
SELECT (SELECT COUNT(*) FROM user) AS users
     , (SELECT COUNT(*) FROM session) AS session
     , (SELECT COUNT(*) FROM game) AS games
     , (SELECT COUNT(*) FROM game where game.play_state = 1) AS games_won
     , (SELECT COUNT(*) FROM game where game.play_state = 2) AS games_lost
     , (SELECT COUNT(*) FROM game where game.play_state = 3) AS games_abandoned
     , (SELECT COUNT(*) FROM game where game.play_state = 4) AS games_current
     , (SELECT max(game.moves) FROM game where game.play_state = 4) AS longest_game
     , (SELECT max(game.score) FROM game where game.play_state = 4) AS highest_score
     , (SELECT CAST(AVG(length(history)*length(history)) - AVG(length(history))*AVG(length(history)) as FLOAT) from game) as history_variance
     , (SELECT avg(length(history)) from game) as history_avg
     , (SELECT max(length(history)) from game) as history_max
     , (SELECT min(length(history)) from game) as history_min
     , (SELECT CAST(total(length(history)) as INT) from game where kind = 2) as history_total
`

type StatsRow struct {
	Users           int64
	Session         int64
	Games           int64
	GamesWon        int64
	GamesLost       int64
	GamesAbandoned  int64
	GamesCurrent    int64
	LongestGame     interface{}
	HighestScore    interface{}
	HistoryVariance interface{}
	HistoryAvg      sql.NullFloat64
	HistoryMax      interface{}
	HistoryMin      interface{}
	HistoryTotal    interface{}
}

func (q *Queries) Stats(ctx context.Context) (StatsRow, error) {
	row := q.db.QueryRowContext(ctx, stats)
	var i StatsRow
	err := row.Scan(
		&i.Users,
		&i.Session,
		&i.Games,
		&i.GamesWon,
		&i.GamesLost,
		&i.GamesAbandoned,
		&i.GamesCurrent,
		&i.LongestGame,
		&i.HighestScore,
		&i.HistoryVariance,
		&i.HistoryAvg,
		&i.HistoryMax,
		&i.HistoryMin,
		&i.HistoryTotal,
	)
	return i, err
}

const updateGame = `-- name: UpdateGame :one
UPDATE game
SET updated_at = ?,
    user_id    = ?,
    rule_id    = ?,
    score      = ?,
    moves      = ?,
    play_state = ?,
    data       = ?,
    history    = ?
WHERE id = ?
RETURNING id, created_at, updated_at, name, description, user_id, rule_id, based_on_game, template_id, score, moves, play_state, data, data_at_start, history
`

type UpdateGameParams struct {
	UpdatedAt sql.NullTime
	UserID    string
	RuleID    string
	Score     int64
	Moves     int64
	PlayState int64
	Data      []byte
	History   []byte
	ID        string
}

func (q *Queries) UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, updateGame,
		arg.UpdatedAt,
		arg.UserID,
		arg.RuleID,
		arg.Score,
		arg.Moves,
		arg.PlayState,
		arg.Data,
		arg.History,
		arg.ID,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.UserID,
		&i.RuleID,
		&i.BasedOnGame,
		&i.TemplateID,
		&i.Score,
		&i.Moves,
		&i.PlayState,
		&i.Data,
		&i.DataAtStart,
		&i.History,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE user
SET updated_at = ?,
    username = ?,
    active_game_id = ?
WHERE id = ?
RETURNING id, created_at, updated_at, username, active_game_id
`

type UpdateUserParams struct {
	UpdatedAt    sql.NullTime
	Username     string
	ActiveGameID string
	ID           string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.UpdatedAt,
		arg.Username,
		arg.ActiveGameID,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.ActiveGameID,
	)
	return i, err
}
