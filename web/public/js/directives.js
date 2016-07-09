var app = angular.module('graphPlotter', []);

app.directive('linePlot', [function () {
    function linkFunc(scope, element, attrs) {
        scope.$watch('graphPlots', function (plots) {
            Plotly.newPlot(element[0], plots);
        });
    }
    return {
        link: linkFunc
    };
}]);
