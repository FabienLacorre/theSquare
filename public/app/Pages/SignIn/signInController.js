'use strict';

var SignIn = angular.module('SignIn', ['ui.router'])

SignIn.config(['$routeProvider', function ($routeProvider) {
    $routeProvider.when('/SignIn', {
        templateUrl: 'Pages/SignIn/signIn.html',
        controller: 'SignInController',
        controllerAs: '$ctrl',
    });
}])

/**
 * @brief SignIn controller 
 */
SignIn.controller('SignInController', function ($location) {

    document.getElementById('test').style.display = "none";

    /**
     * @brief Cancel button handler
     */
    this.cancel = () => {
        console.log("cancel account")
        $location.path('/');
    }

    /**
     * @brief validation button handler
     */
    this.validate = () => {
        console.log("validation account")
        $location.path('/Profile');
    }
});