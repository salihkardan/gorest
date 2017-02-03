var app = angular.module("MyApp", [
	'ui.router',
	'ui.bootstrap',
	'chart.js',
	'graphPlotter',
	'tc.chartjs'
]);

app.config(function($stateProvider, $urlRouterProvider, $httpProvider, $logProvider, ChartJsProvider) {
	$logProvider.debugEnabled(false);

	ChartJsProvider.setOptions({
		 colours: ['#97BBCD', '#DCDCDC', '#F7464A', '#46BFBD', '#FDB45C', '#949FB1', '#4D5360'],
		 responsive: true
	 });

	// Configure all doughnut charts
	ChartJsProvider.setOptions('Doughnut', {
		 animateScale: true
	 });

	// Routes and states
	$urlRouterProvider.otherwise("/events");
	$stateProvider
		.state('home', {
			url: "/home",
			templateUrl: "partials/home.html",
		})
		.state('events', {
			url: "/events",
			templateUrl: "partials/events.html",
			controller: "EventsController",
		})
		.state('requests', {
			url: "/requests",
			templateUrl: "partials/requests.html",
			controller: "RequestController",
		})
		.state('monitor', {
			url: "/monitor",
			templateUrl: "partials/monitor.html",
			controller: "MonitorController",
		})
	// HTTP Interceptor for config
	$httpProvider.interceptors.push('APIInterceptor');
});

app.run(function($rootScope, $location, $state) {
	// Check if logged in
	if ( localStorage.token ){
		$rootScope.token = localStorage.token;
	}

	$rootScope.lang = "en";
	$rootScope.changeLang = function(key) {
		$rootScope.lang = key;
		$translate.use(key);
	};

	$rootScope.logout = function() {
		// TODO: Remove Token
		delete localStorage.token;
		$rootScope.token = null;
		$state.go("events");
	}

	$rootScope.$on('$stateChangeStart',
		function(event, toState, toParams, fromState, fromParams) {
			// Check tostate
      var whitelist  = ["signup", "login"];
			if (!$rootScope.token) {
				if ( !whitelist.indexOf(toState)){
					$state.go("events")
					event.preventDefault();
				}
			}
		});
})

app.factory('APIInterceptor', function($q, $rootScope) {
	return {
		// optional method
		'request': function(config) {
			// TODO: Check paths for api endpoints
			// TODO: Check if login , not redirect to login
			// TODO: Don't check token on login
			// TODO: Not all 401s mean login required
			if ($rootScope.token) {
				config.headers['Authorization'] = 'Bearer ' + $rootScope.token;
			}
			return config;
		},

		// optional method
		'requestError': function(rejection) {
			// do something on error
			return $q.reject(rejection);
		},

		// optional method
		'response': function(response) {
			// do something on success
			return response;
		},

		// optional method
		'responseError': function(rejection) {
			// do something on error
			return $q.reject(rejection);
		}
	};
});
