package db

import (
	"time"

	"github.com/plandex/plandex/shared"
)

// The models below should only be used server-side.
// Many of them have corresponding models in shared/api for client-side use.
// This adds some duplication, but helps ensure that server-only data doesn't leak to the client.
// Models used client-side have a ToApi() method to convert it to the corresponding client-side model.

type Org struct {
	Id                 string    `db:"id"`
	Name               string    `db:"name"`
	Domain             string    `db:"domain"`
	AutoAddDomainUsers bool      `db:"auto_add_domain_users"`
	CreatorId          string    `db:"creator_id"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

func (org *Org) ToApi() *shared.Org {
	return &shared.Org{
		Id:   org.Id,
		Name: org.Name,
	}
}

type User struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (user *User) ToApi() *shared.User {
	return &shared.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

type OrgUser struct {
	Id        string    `db:"id"`
	OrgId     string    `db:"org_id"`
	UserId    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Project struct {
	Id               string    `db:"id"`
	OrgId            string    `db:"org_id"`
	Name             string    `db:"name"`
	LastActivePlanId string    `db:"last_active_plan_id"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

func (project *Project) ToApi() *shared.Project {
	return &shared.Project{
		Id:               project.Id,
		Name:             project.Name,
		LastActivePlanId: project.LastActivePlanId,
	}
}

type UserProject struct {
	Id               string    `db:"id"`
	OrgId            string    `db:"org_id"`
	UserId           string    `db:"user_id"`
	ProjectId        string    `db:"project_id"`
	LastActivePlanId string    `db:"last_active_plan_id"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

type Plan struct {
	Id            string            `db:"id"`
	OrgId         string            `db:"org_id"`
	CreatorId     string            `db:"creator_id"`
	ProjectId     string            `db:"project_id"`
	Name          string            `db:"name"`
	Status        shared.PlanStatus `db:"status"`
	Error         string            `db:"error"`
	ContextTokens int               `db:"context_tokens"`
	ConvoTokens   int               `db:"convo_tokens"`
	AppliedAt     *time.Time        `db:"applied_at,omitempty"`
	ArchivedAt    *time.Time        `db:"archived_at,omitempty"`
	CreatedAt     time.Time         `db:"created_at"`
	UpdatedAt     time.Time         `db:"updated_at"`
}

func (plan *Plan) ToApi() *shared.Plan {
	return &shared.Plan{
		Id:            plan.Id,
		CreatorId:     plan.CreatorId,
		Name:          plan.Name,
		Status:        plan.Status,
		ContextTokens: plan.ContextTokens,
		ConvoTokens:   plan.ConvoTokens,
		AppliedAt:     plan.AppliedAt,
		ArchivedAt:    plan.ArchivedAt,
		CreatedAt:     plan.CreatedAt,
		UpdatedAt:     plan.UpdatedAt,
	}
}

type ConvoSummary struct {
	Id                          string    `db:"id"`
	OrgId                       string    `db:"org_id"`
	PlanId                      string    `db:"plan_id"`
	LatestConvoMessageId        string    `db:"latest_convo_message_id"`
	LatestConvoMessageCreatedAt time.Time `db:"latest_convo_message_created_at"`
	Summary                     string    `db:"summary"`
	Tokens                      int       `db:"tokens"`
	NumMessages                 int       `db:"num_messages"`
	CreatedAt                   time.Time `db:"created_at"`
}

func (summary *ConvoSummary) ToApi() *shared.ConvoSummary {
	return &shared.ConvoSummary{
		Id:                          summary.Id,
		LatestConvoMessageId:        summary.LatestConvoMessageId,
		LatestConvoMessageCreatedAt: summary.LatestConvoMessageCreatedAt,
		Summary:                     summary.Summary,
		Tokens:                      summary.Tokens,
		NumMessages:                 summary.NumMessages,
		CreatedAt:                   summary.CreatedAt,
	}
}

type ConvoMessageDescription struct {
	Id                    string    `db:"id"`
	OrgId                 string    `db:"org_id"`
	PlanId                string    `db:"plan_id"`
	ConvoMessageId        string    `db:"convo_message_id"`
	SummarizedToMessageId string    `db:"summarized_to_message_id"`
	MadePlan              bool      `db:"made_plan"`
	CommitMsg             string    `db:"commit_msg"`
	Files                 []string  `db:"files"`
	Error                 string    `db:"error"`
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
}

func (desc *ConvoMessageDescription) ToApi() *shared.ConvoMessageDescription {
	return &shared.ConvoMessageDescription{
		Id:                    desc.Id,
		ConvoMessageId:        desc.ConvoMessageId,
		SummarizedToMessageId: desc.SummarizedToMessageId,
		MadePlan:              desc.MadePlan,
		CommitMsg:             desc.CommitMsg,
		Files:                 desc.Files,
		Error:                 desc.Error,
		CreatedAt:             desc.CreatedAt,
		UpdatedAt:             desc.UpdatedAt,
	}
}

type PlanBuild struct {
	Id             string    `db:"id"`
	OrgId          string    `db:"org_id"`
	PlanId         string    `db:"plan_id"`
	ConvoMessageId string    `db:"convo_message_id"`
	Error          string    `db:"error"`
	ErrorPath      string    `db:"error_path"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (build *PlanBuild) ToApi() *shared.PlanBuild {
	return &shared.PlanBuild{
		Id:             build.Id,
		ConvoMessageId: build.ConvoMessageId,
		Error:          build.Error,
		ErrorPath:      build.ErrorPath,
		CreatedAt:      build.CreatedAt,
		UpdatedAt:      build.UpdatedAt,
	}
}

// Models below are stored in files, not in the database.
// This allows us to store them in a git repo and use git to manage history.

type Context struct {
	Id          string             `json:"id"`
	OrgId       string             `json:"orgId"`
	CreatorId   string             `json:"creatorId"`
	PlanId      string             `json:"planId"`
	ContextType shared.ContextType `json:"contextType"`
	Name        string             `json:"name"`
	Url         string             `json:"url"`
	FilePath    string             `json:"filePath"`
	Sha         string             `json:"sha"`
	NumTokens   int                `json:"numTokens"`
	Body        string             `json:"body,omitempty"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

func (context *Context) ToApi() *shared.Context {
	return &shared.Context{
		Id:          context.Id,
		CreatorId:   context.CreatorId,
		ContextType: context.ContextType,
		Name:        context.Name,
		Url:         context.Url,
		FilePath:    context.FilePath,
		Sha:         context.Sha,
		NumTokens:   context.NumTokens,
		Body:        context.Body,
		CreatedAt:   context.CreatedAt,
		UpdatedAt:   context.UpdatedAt,
	}
}

type ConvoMessage struct {
	Id        string    `json:"id"`
	OrgId     string    `json:"orgId"`
	PlanId    string    `json:"planId"`
	UserId    string    `json:"userId"`
	Role      string    `json:"role"`
	Tokens    int       `json:"tokens"`
	Num       int       `json:"num"`
	Message   string    `json:"message"`
	Stopped   bool      `json:"stopped"`
	CreatedAt time.Time `json:"createdAt"`
}

func (msg *ConvoMessage) ToApi() *shared.ConvoMessage {
	return &shared.ConvoMessage{
		Id:        msg.Id,
		UserId:    msg.UserId,
		Role:      msg.Role,
		Tokens:    msg.Tokens,
		Num:       msg.Num,
		Message:   msg.Message,
		Stopped:   msg.Stopped,
		CreatedAt: msg.CreatedAt,
	}
}

type PlanFileResult struct {
	Id           string                `json:"id"`
	OrgId        string                `json:"orgId"`
	PlanId       string                `json:"planId"`
	PlanBuildId  string                `json:"planBuildId"`
	Path         string                `json:"path"`
	ContextSha   string                `json:"contextSha"`
	Content      string                `json:"content,omitempty"`
	Replacements []*shared.Replacement `json:"replacements"`
	AnyFailed    bool                  `json:"anyFailed"`
	Error        string                `json:"error"`
	AppliedAt    *time.Time            `json:"appliedAt,omitempty"`
	RejectedAt   *time.Time            `json:"rejectedAt,omitempty"`
	CreatedAt    time.Time             `json:"createdAt"`
	UpdatedAt    time.Time             `json:"updatedAt"`
}

func (res *PlanFileResult) ToApi() *shared.PlanFileResult {
	return &shared.PlanFileResult{
		Id:           res.Id,
		PlanBuildId:  res.PlanBuildId,
		Path:         res.Path,
		ContextSha:   res.ContextSha,
		Content:      res.Content,
		AnyFailed:    res.AnyFailed,
		AppliedAt:    res.AppliedAt,
		RejectedAt:   res.RejectedAt,
		Replacements: res.Replacements,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}
}
