'use strict';

var Connection = angular.module('Connection', ['ui.router'])

Connection.config(['$routeProvider', function ($routeProvider) {
    $routeProvider.when('/', {
        templateUrl: 'Pages/Connection/connection.html',
        controller: 'ConnectionController',
        controllerAs: '$ctrl',
    });
}])

/**
 * @brief Connection controller 
 */
Connection.controller('ConnectionController', function ($location, $http) {

    document.getElementById('test').style.display = "none";

    /**
     * @brief try to log a user 
     * if OK redirect to dashboard page
     */
    this.tryToLogin = () => {
        // TO DO connection routine
        $http.post('/api/login', null, { headers: { 'Authorization': 'Basic dG90bzp0YXRh' } }).then((response) => response.data)
            .then((response) => {
                $location.path('/Dashboard');
            }).catch((error) => console.error(error))
    }

    /**
     * @brief redirect user to sign in page
     */
    this.goSignIn = () => {
        console.log("SIGN IN ")
        $location.path('/SignIn');
    }
});