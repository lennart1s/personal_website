import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';

import { Auth0Plugin } from './auth';

Vue.use(Auth0Plugin, {
  domain: 'sandbothe-dev.eu.auth0.com',
  clientId: 'S7FcgoUvto1cK8sFqtAkyUHXUMdPmZsp',
  /* eslint-disable @typescript-eslint/no-explicit-any */
  onRedirectCallback: (appState: any) => {
    router.push(appState && appState.targetUrl ? appState.targetUrl : window.location.pathname);
  },
  // redirectUri: 'http://localhost:9090/api/auth/callback',
});

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
