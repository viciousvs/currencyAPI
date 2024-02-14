package config

const cronSpec = "JobCronTime"

type CronConfig struct {
	Spec int
}

func MakeCronConfig() CronConfig {
	return CronConfig{
		Spec: getEnvAsInt(cronSpec, 1),
	}
}
