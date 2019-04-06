'use strict';

var Profile = angular.module('Proposition', ['ngRoute'])

Profile.config(['$routeProvider', function ($routeProvider) {
  $routeProvider.when('/Proposition', {
    templateUrl: 'Pages/Proposition/proposition.html',
    controller: 'PropositionCtrl',
    controllerAs: '$ctrl',
  });
}])

  /**
  * @brief Proposals controller
  */
  Profile.controller('PropositionCtrl', function ($location, $http, $scope) {
    console.log("hello proposition controller")
    this.userId = localStorage.getItem("id");

    this.userConnected = undefined
    $http.get('/api/profile/'+ this.userId+ '/propositions/users').then((response) => response.data)
    .then((response) => {
      this.userConnected = response
      console.log("users", response)
    }).catch(() => alert("Error loading datas"))

    this.hobbiesConnected = undefined
    $http.get('/api/profile/'+ this.userId+ '/propositions/users/hobbies').then((response) => response.data)
    .then((response) => {
      this.hobbiesConnected = response
      console.log("hobbies", response)
    }).catch(() => alert("Error loading datas"))

    this.companiesConnected = undefined
    $http.get('/api/profile/'+ this.userId+ '/propositions/companies').then((response) => response.data)
    .then((response) => {
      this.companiesConnected = response
      console.log("companies", response)
    }).catch(() => alert("Error loading datas"))

  }); 


