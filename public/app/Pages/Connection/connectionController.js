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
    this.username = undefined;
    this.password = undefined;


    //login:mdp Base64
    /**
     * @brief try to log a user 
     * if OK redirect to dashboard page
     */
    this.tryToLogin = () => {
        console.log(this.username)
        console.log(this.password)
        let base64 = btoa(this.username + ":" + this.password);
        console.log(base64)
        $http.post('/api/login', null, { headers: { Authorization: 'Basic ' + base64 } }).then((response) => response.data)
            .then((response) => {
                console.log(response)
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