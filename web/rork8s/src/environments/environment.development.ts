declare let require: any;
const { version: appVersion } = require('../../package.json');

export const environment = {
  appVersion,
  configPath: 'assets/config/config.json',
  production: false,
};
