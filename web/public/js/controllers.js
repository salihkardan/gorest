var app = angular.module("MyApp");

app.controller("EventsController", function($scope, $http) {
	$scope.events = [];
	$http.get("/api/events").then(function success(resp) {
		$scope.events = resp.data;
	});
});

app.controller("MonitorController", function($scope, $http) {

	function timeConverter(UNIX_timestamp){
	  var a = new Date(UNIX_timestamp);
	  var months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
	  var year = a.getFullYear();
	  var month = months[a.getMonth()];
	  var date = a.getDate();
	  var hour = a.getHours();
	  var min = a.getMinutes();
	  var sec = a.getSeconds();
	  // var time = date + ' ' + month + ' ' + year + ' ' + hour + ':' + min + ':' + sec ;
		var time = hour + ':' + min + ':' + sec ;
	  return time;
	}

	$scope.resp = [];
	var myArray = [];
	var timeArray = [];

	var myArray2 = [];
	var timeArray2 = [];


	$http.get("/api/responses").then(function success(resp) {
		$scope.resp = resp.data;
	  for(var j=0; j<$scope.resp.length; j++) {
				// console.log($scope.resp[j])
				if ( $scope.resp[j].APIKey == "test-api-key"){
					myArray2.push($scope.resp[j].Duration);
					timeArray2.push(timeConverter($scope.resp[j].Timestamp))
				} else {
					myArray.push($scope.resp[j].Duration);
					timeArray.push(timeConverter($scope.resp[j].Timestamp))
				}

	  }
		var trace1 = {
			x: timeArray,
			y: myArray,
				type: 'scatter',
				name: "def"
		};

		var trace2 = {
				x: timeArray2,
				y: myArray2,
				type: 'scatter',
				name: "abc"
		};
		$scope.graphPlots = [trace1, trace2];
	});

});
