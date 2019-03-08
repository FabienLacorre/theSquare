'use strict';

var Details = angular.module('Details', ['ngRoute'])

Details.config(['$routeProvider', function ($routeProvider) {
  $routeProvider.when('/Details/:type', {
    templateUrl: 'Pages/Details/details.html',
    controller: 'DetailsController',
    controllerAs: '$ctrl',
  });
}])

Details.controller('DetailsController', function ($routeParams, $location) {
    console.log($routeParams.type)
    this.type = $routeParams.type;

    this.object = {
        name: "TEST NAME",
        description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        date: new Date(),
        type: this.type,
        photo: "../../img/test.jpg",
    }

    this.changeLocation = (route) => {
        $location.path('/' + route)
    }

})