package model

type CronJob struct {
	ApiVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Metadata   Metadata    `yaml:"metadata"`
	Spec       CronJobSpec `yaml:"spec"`
}

type CronJobSpec struct {
	// Schedule is defined with cron format
	Schedule string `yaml:"schedule"`
	// ConcurrencyPolicy should be either "Allow", "Forbid", or "Replace"
	ConcurrencyPolicy string `yaml:"concurrencyPolicy"`
	// Suspend should be either "true" or "false"
	Suspend     string              `yaml:"suspend"`
	JobTemplate CronJobSpecTemplate `yaml:"jobTemplate"`
}

type CronJobSpecTemplate struct {
	Metadata Metadata       `yaml:"metadata"`
	Spec     DeploymentSpec `yaml:"spec"`
}
