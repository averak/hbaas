package main

import "github.com/averak/hbaas/entrypoint/batch_job/job"

var registry = map[string]job.BatchJob{
	"purge_old_echos": job.NewPurgeOldEchos(),
}
