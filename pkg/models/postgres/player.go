package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/alexwilkerson/ddstats-api/pkg/models"
	"github.com/jmoiron/sqlx"
)

// PlayerModel wraps the database connection
type PlayerModel struct {
	DB *sqlx.DB
}

func (p *PlayerModel) Insert(player *models.Player) error {
	stmt := `INSERT INTO player(
			player_name,
			rank,
			game_time,
			death_type,
			gems,
			daggers_hit,
			daggers_fired,
			enemies_killed,
			accuracy,
			overall_time,
			overall_deaths,
			overall_gems,
			overall_enemies_killed,
			overall_daggers_hit,
			overall_daggers_fired,
			overall_accuracy
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`
	_, err := p.DB.Exec(stmt,
		player.PlayerName,
		player.Rank,
		player.GameTime,
		player.DeathType,
		player.Gems,
		player.DaggersHit,
		player.DaggersFired,
		player.EnemiesKilled,
		player.Accuracy,
		player.OverallTime,
		player.OverallDeaths,
		player.OverallGems,
		player.OverallEnemiesKilled,
		player.OverallDaggersHit,
		player.OverallDaggersFired,
		player.OverallAccuracy,
	)
	if err != nil {
		return err
	}

	return nil
}

// Get returns a single player record
func (p *PlayerModel) Get(id int) (*models.Player, error) {
	var player models.Player

	stmt := "SELECT * FROM player WHERE id=$1"
	err := p.DB.Get(&player, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return &player, nil
}

// GetAll retreives a slice of users using a specified page size and page num starting at 1
func (p *PlayerModel) GetAll(pageSize, pageNum int) ([]*models.Player, error) {
	var players []*models.Player

	stmt := fmt.Sprintf("SELECT * FROM player WHERE id<>-1 ORDER BY game_time DESC LIMIT %d OFFSET %d", pageSize, (pageNum-1)*pageSize)
	err := p.DB.Select(&players, stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return players, nil
}

// GetTotalCount returns the total number of players in the database
func (p *PlayerModel) GetTotalCount() (int, error) {
	var playerCount int
	stmt := "SELECT COUNT(1) FROM player WHERE id<>-1"
	err := p.DB.QueryRow(stmt).Scan(&playerCount)
	if err != nil {
		return 0, err
	}
	return playerCount, nil
}
