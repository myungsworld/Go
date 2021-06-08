package _map

import (
	"git-codecommit.ap-northeast-2.amazonaws.com/v1/repos/boss-api/backend/middlewares"
	"git-codecommit.ap-northeast-2.amazonaws.com/v1/repos/boss-api/backend/models"
	"net/http"
)

func MapHandling() {
	nCompetitionId := middlewares.GetNCompetitionId(c)
	eventId := middlewares.GetNEventId(c)
	teacherId := middlewares.GetReqTeacherIdFromToken(c.Request)
	teacher := middlewares.FetchTeacherById(teacherId, "school_info_id")

	var body ReqInfo
	if err := c.ShouldBind(&body); err != nil {
		panic(cerror.BadRequest())
	}

	type raw struct {
		TeamId    int    `json:"team_id"`
		Detail    string `json:"detail"`
		Result    string `json:"result"`
		VideoId   int    `json:"video_id"`
		Grade     string `json:"grade"`
		ClassName string `json:"class_name"`
		Number    int    `json:"number"`
		Name      string `json:"name"`
	}

	fetchTeamUser := datastore.DB.Model(&models.Nto1Team{}).
		Select([]string{
			"nto1_teams.id as team_id",
			"nto1_teams.detail as detail",
			"u.grade as grade",
			"u.class_name as class_name",
			"u.number as number",
			"u.name as name",
			"nto1_teams.result",
			"nto1_teams.video_id",
		}).
		Joins("join nto1_team_users ntu on ntu.nto1_team_id = nto1_teams.id").
		Joins("join users u on ntu.user_id = u.id").
		Where("nto1_teams.competition_id = ?", nCompetitionId).
		Where("nto1_teams.event_id = ?", eventId).
		Where("nto1_teams.school_info_id = ?", teacher.SchoolInfoId).
		Order("nto1_teams.id asc")

	if body.Detail != "" {
		fetchTeamUser = fetchTeamUser.Where("detail = ?", body.Detail)
	}

	if body.Name != "" {
		fetchTeamUser = fetchTeamUser.Where(fmt.Sprintf("u.name LIKE '%%%s%%'", body.Name))
	}

	var totalCnt int64
	if err := fetchTeamUser.
		Count(&totalCnt).Error; err != nil {
		panic(cerror.DBErr(err))
	}
	event := middlewares.FetchNationalCompetitionEventById(eventId, "team_count")
	totalCnt = totalCnt / int64(event.TeamCount)

	users := make([]raw, 0)
	if err := fetchTeamUser.
		Offset((body.PageIndex - 1) * pageSize).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		panic(cerror.DBErr(err))
	}

	teamIdMap := make(map[int][]raw)
	for _, user := range users {
		if _, exist := teamIdMap[user.TeamId]; !exist {
			teamIdMap[user.TeamId] = make([]raw, 0)
		}
		teamIdMap[user.TeamId] = append(teamIdMap[user.TeamId],user)
	}

	teams := make([]RespTeam,0)
	for teamId, els := range teamIdMap {
		team := RespTeam{
			TeamId: teamId,
			Detail: els[0].Detail,
			VideoId: els[0].VideoId,
			Result: els[0].Result,
		}
		team.Users = make([]models.User,0)
		for _ , el := range els {
			team.Users = append(team.Users, models.User{
				Name: el.Name,
				Grade: el.Grade,
				ClassName: el.ClassName,
				Number: el.Number,
			})
		}
		teams = append(teams, team)
	}

	c.JSON(http.StatusOK, models.PageInfo{
		Total:     totalCnt,
		PageIndex: body.PageIndex,
		PageSize:  pageSize,
		Contents:  teams,
	})

}
