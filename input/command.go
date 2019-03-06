package input

import (
	"log"
	"strings"
)

// ParseCommand is doing adhoc command parsing.
// for the future, we should write an actual parser.
func ParseCommand(raw string) (ok bool, cmd interface{}) {
	log.Printf("debug: input: %v\n", raw)
	tmp := strings.Split(raw, "\n")

	// If there are no possibility that the comment body is not formatted
	// `@botname command`, stop to process.
	if len(tmp) < 1 {
		return false, nil
	}

	for index := range tmp {
		if strings.Index(tmp[index], "@") != -1 {
			body := tmp[index]
			log.Printf("debug: body: %v\n", body)

			r := strings.NewReader(body)
			p := newParser(r)
			cmd, err := p.Parse()
			if err != nil {
				log.Printf("debug: parse error: %v\n", err)
				return false, nil
			}
			return true, cmd
		}
	}

	return false, nil
}

type AcceptChangesetCommand interface {
	BotName() string
}

type AcceptChangeByReviewerCommand struct {
	botName string
}

func (s *AcceptChangeByReviewerCommand) BotName() string {
	return s.botName
}

type AcceptChangeByOthersCommand struct {
	botName  string
	Reviewer []string
}

func (s *AcceptChangeByOthersCommand) BotName() string {
	return s.botName
}

type AcceptChangeOnReview struct {
	BotNameForReview string
}

func (s *AcceptChangeOnReview) BotName() string {
	return s.BotNameForReview
}

type AssignReviewerCommand struct {
	Reviewer []string
}

type CancelApprovedByReviewerCommand struct {
	botName string
}

func (s *CancelApprovedByReviewerCommand) BotName() string {
	return s.botName
}
