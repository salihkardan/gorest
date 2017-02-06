var app = angular.module("MyApp");

app.controller("EventsController", function ($scope, $http) {
    $scope.events = [];
    $http.get("/api/events").then(function success(resp) {
        $scope.events = resp.data;
    });
});

app.controller("RequestController", function ($scope, $http) {
    $scope.requests = [];
    $http.get("/api/requests").then(function success(resp) {
        $scope.requests = resp.data;
        console.log(resp.data)
    });
});

app.controller("HomeController", function ($scope, $http, $location) {
    $http.get('http://ipinfo.io/json').success(
        function (data) {
            $http.post("/incoming", {
                ip: data.ip,
                country: data.country,
                city: data.city
            }).then(function success(resp) {
                console.log(resp.data)
            });
        });
    $location.html5Mode(true)
});


app.controller("MonitorController", function ($scope, $http) {
    // todo: may be a lot better
    function count(respArray) {
        var a = 0, b = 0, c = 0, d = 0, e = 0, f = 0;
        var result = [];
        for (var j = 0; j < respArray.length; j++) {
            if (respArray[j].Duration < 1) {
                a++;
            } else if (respArray[j].Duration > 1 && respArray[j].Duration < 5) {
                b++;
            } else if (respArray[j].Duration > 5 && respArray[j].Duration < 10) {
                c++;
            } else if (respArray[j].Duration > 10 && respArray[j].Duration < 20) {
                d++;
            } else if (respArray[j].Duration > 20 && respArray[j].Duration < 50) {
                e++;
            } else if (respArray[j].Duration > 50) {
                f++;
            }
        }
        result.push(a, b, c, d, e, f);
        return result;
    }

    $http.get("/api/requests").then(function success(resp) {
        var dataArray = [];
        if (resp.data != null) {
            $scope.resp = resp.data;
            dataArray = count($scope.resp)
        }

        // Chart.js Data
        $scope.data = {
            labels: ['<1ms', '1-5ms', '5-10ms', '10-20ms', '20-50ms', '50-100ms'],
            datasets: [
                {
                    label: 'Number of requests',
                    fillColor: 'rgba(220,220,220,0.5)',
                    strokeColor: 'rgba(220,220,220,0.8)',
                    highlightFill: 'rgba(220,220,220,0.75)',
                    highlightStroke: 'rgba(220,220,220,1)',
                    data: dataArray
                }
            ]
        };
    });
});
