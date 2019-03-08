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
Connection.controller('ConnectionController', function ($location) {

    document.getElementById('test').style.display = "none";

    /**
     * @brief try to log a user 
     * if OK redirect to dashboard page
     */
    this.tryToLogin = () => {
        // TO DO connection routine
        $location.path('/Dashboard');
    }

    /**
     * @brief redirect user to sign in page
     */
    this.goSignIn = () => {
        console.log("SIGN IN ")
        $location.path('/SignIn');
    }
});