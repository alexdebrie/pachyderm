package server

import (
	"sort"

	"go.pedge.io/proto/time"

	"github.com/pachyderm/pachyderm/src/pps/persist"
)

// TODO: this should be a call through the actual persist storage
//
// This does not work:
//
//     func(term gorethink.Term) gorethink.Term {
//         return term.OrderBy(gorethink.Desc("created_at"))
//     }

func sortJobInfosByTimestampDesc(s []*persist.JobInfo) {
	sort.Sort(jobInfosByTimestampDesc(s))
}

type jobInfosByTimestampDesc []*persist.JobInfo

func (s jobInfosByTimestampDesc) Len() int          { return len(s) }
func (s jobInfosByTimestampDesc) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }
func (s jobInfosByTimestampDesc) Less(i int, j int) bool {
	return prototime.TimestampLess(s[j].CreatedAt, s[i].CreatedAt)
}

func sortJobStatusesByTimestampDesc(s []*persist.JobStatus) {
	sort.Sort(jobStatusesByTimestampDesc(s))
}

type jobStatusesByTimestampDesc []*persist.JobStatus

func (s jobStatusesByTimestampDesc) Len() int          { return len(s) }
func (s jobStatusesByTimestampDesc) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }
func (s jobStatusesByTimestampDesc) Less(i int, j int) bool {
	return prototime.TimestampLess(s[j].Timestamp, s[i].Timestamp)
}

func sortJobLogsByTimestampAsc(s []*persist.JobLog) {
	sort.Sort(jobLogsByTimestampAsc(s))
}

type jobLogsByTimestampAsc []*persist.JobLog

func (s jobLogsByTimestampAsc) Len() int          { return len(s) }
func (s jobLogsByTimestampAsc) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }
func (s jobLogsByTimestampAsc) Less(i int, j int) bool {
	return prototime.TimestampLess(s[i].Timestamp, s[j].Timestamp)
}

func sortPipelineInfosByTimestampDesc(s []*persist.PipelineInfo) {
	sort.Sort(pipelineInfosByTimestampDesc(s))
}

type pipelineInfosByTimestampDesc []*persist.PipelineInfo

func (s pipelineInfosByTimestampDesc) Len() int          { return len(s) }
func (s pipelineInfosByTimestampDesc) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }
func (s pipelineInfosByTimestampDesc) Less(i int, j int) bool {
	return prototime.TimestampLess(s[j].CreatedAt, s[i].CreatedAt)
}
