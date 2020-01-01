# ddstats-api

### TODO: endpoints

#### index

- [ ] GET user/live
- [ ] GET game/top
- [ ] GET game/recent?pagesize={int}&pagenum={int}
- [ ] GET client/version/latest
- [ ] GET news/recent?num={int}

#### games

- [ ] GET game/recent?pagesize={int}&pagenum={int}

#### game_log

- [ ] GET game/info?id={int}
- [ ] GET game?id={int} (all)
- [ ] GET game/gems?id={int}
- [ ] GET game/homing-daggers?id={int}
- [ ] GET game/accuracy?id={int}
- [ ] GET game/enemies-alive?id={int}
- [ ] GET game/enemies-killed?id={int}

#### users

- [ ] GET user?pagesize={int}&pagenum={int} (list all users)

#### dd backend

- [x] GET api/v2/ddapi/get_user_by_rank?rank={int}
- [x] GET api/v2/ddapi/get_user_by_id?id={int}
- [x] GET api/v2/ddapi/user_search?user={string}
- [x] GET api/v2/ddapi/get_scores?offset={int}&limit={int}

### importing csvs from app.db

In order to port the database from SQLite3 to Postgres, the following shenanigans must occur:

- scp app.db from server, then: `sqlite3 app.db`
- Afterwards, import the csvs via Postico.
- See `schema.sql` for creating the Postgres database.

#### game

```sql
.headers on
.mode csv
.output state.csv
-- printf required because of sqlite's weird float formatting
select id, player_id, granularity, printf("%.6f", game_time) as game_time, death_type, gems, homing_daggers, daggers_fired, daggers_hit, enemies_alive, enemies_killed, time_stamp, replay_player_id, survival_hash, version, printf("%.6f", level_two_time) as level_two_time, printf("%.6f", level_three_time) as level_three_time, printf("%.6f", level_four_time) as level_four_time, printf("%.6f", homing_daggers_max_time) as homing_daggers_max_time, printf("%.6f", enemies_alive_max_time) as enemies_alive_max_time, homing_daggers_max, enemies_alive_max from game;
```

#### state

```sql
.headers on
.mode csv
.output state.csv
-- remove any inconsistencies between state and game tables
delete from state where game_id not in (select id from game where id is not null);
-- printf required because of sqlite's weird float formatting
select id, game_id, printf("%.6f", game_time) as game_time, gems, homing_daggers, daggers_hit, daggers_fired, enemies_alive, enemies_killed from state;
```

#### spawnset

```sql
.headers on
.mode csv
.output state.csv
select * from spawnset;
```

#### player (renamed from user)

```sql
.headers on
.mode csv
.output player.csv
select * from user;
```

#### live

```sql
.headers on
.mode csv
.output live.csv
select * from live;
```
