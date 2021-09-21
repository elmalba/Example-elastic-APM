import { init as initApm } from '@elastic/apm-rum'

const apm = initApm({

  // Set required service name (allowed characters: a-z, A-Z, 0-9, -, _, and space)
  serviceName: 'frontend',

  // Set custom APM Server URL (default: http://localhost:8200)
  serverUrl: 'https://54f4fdeaa1bf4d52b857c96ca785a140.apm.eastus2.azure.elastic-cloud.com:443',

  // Set service version (required for sourcemap feature)
  serviceVersion: '0.1'
})


export default apm 