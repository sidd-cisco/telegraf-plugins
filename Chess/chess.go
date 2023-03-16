package chess

// chess.go

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type ResponseLeaderboards struct {
	PlayerID int    `json:"player_id"`
	Username string `json:"username"`
	Rank     int    `json:"rank"`
	Score    int    `json:"score"`
}

type Leaderboards struct {
	Daily []ResponseLeaderboards `json:"daily"`
}

type Chess struct {
	profiles    []string        `toml:"ok"`
	leaderboard bool            `toml:"ok"`
	streamers   bool            `toml:"ok"`
	Countries   []string        `toml:"ok"`
	Log         telegraf.Logger `toml:"-"`
}

func (c *Chess) Description() string {
	return "a chess plugin"
}

func (c *Chess) SampleConfig() string {
	return `
 ## Indicate if everything is fine
 ok = true
`
}

// Init is for setup, and validating config.
func (c *Chess) Init() error {
	return nil
}

func (c *Chess) Gather(acc telegraf.Accumulator) error {
	if c.leaderboard {
		// Obtain all public leaderboard information from the
		// chess.com api

		var leaderboards Leaderboards
		// request and unmarshall leaderboard information
		// and add it to the accumulator
		resp, err := http.Get("https://api.chess.com/pub/leaderboards")
		if err != nil {
			c.Log.Errorf("failed to GET leaderboards json: %w", err)
			return err
		}

		data, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			c.Log.Errorf("failed to read leaderboards json response body: %w", err)
			return err
		}

		//unmarshall the data
		err = json.Unmarshal(data, &leaderboards)
		if err != nil {
			c.Log.Errorf("failed to unmarshall leaderboards json: %w", err)
			return err
		}

		for _, stat := range leaderboards.Daily {
			var fields = make(map[string]interface{}, len(leaderboards.Daily))
			var tags = map[string]string{
				"playerId": strconv.Itoa(stat.PlayerID),
			}
			fields["username"] = stat.Username
			fields["rank"] = stat.Rank
			fields["score"] = stat.Score
			acc.AddFields("leaderboards", fields, tags)
		}
	}
	return nil
}

func init() {
	inputs.Add("chess", func() telegraf.Input { return &Chess{} })
}
