import { init as initApm } from '@elastic/apm-rum'

const apm = initApm({

  // Set required service name (allowed characters: a-z, A-Z, 0-9, -, _, and space)
  serviceName: 'frontend',

  // Set custom APM Server URL (default: http://localhost:8200)
  serverUrl: 'https://7f7d8ec9b82e482a88dd64e7d57f8e7c.apm.us-west1.gcp.cloud.es.io:443',

  // Set service version (required for sourcemap feature)
  serviceVersion: '0.1'
})


export default apm 