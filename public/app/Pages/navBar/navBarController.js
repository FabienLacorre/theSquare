var navBar = angular.module('navBar', [])
navBar.controller('navBarController', function ($location, $http, $scope) {
    console.log("THIS IS A TEST NAV BAR")
    
    $scope.logout = () => {
        $http.post('/api/logout').then((response) => response.data)
        .then((response) => {
            console.log(response)
            $location.path("#!/");
        }).catch((err) => console.error(err))
    }
});

// neo4j-admin load --from=data/preset/2019-03-11.dump --force