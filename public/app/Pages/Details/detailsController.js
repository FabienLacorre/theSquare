'use strict';

var Details = angular.module('Details', ['ngRoute'])

Details.config(['$routeProvider', function ($routeProvider) {
  $routeProvider.when('/Details/:type/:id', {
    templateUrl: 'Pages/Details/details.html',
    controller: 'DetailsController',
    controllerAs: '$ctrl',
  });
}])

Details.controller('DetailsController', function ($routeParams, $location, $http, $scope) {
  console.log($routeParams.type)
  this.type = $routeParams.type;
  this.id = $routeParams.id;

  console.log(this.id)
  const promise = []
  if (this.type === "companie") {
    promise.push($http.get('/api/company/' + this.id))
  } else if (this.type === "friend") {
    console.log("friend")
    promise.push($http.get('/api/profile/' + this.id))
  } else if (this.type === "job") {
    promise.push($http.get('/api/job/' + this.id))
  } else if (this.type === "hobbie") {
    promise.push($http.get('/api/hobby/' + this.id))
  } else if (this.type === "skill") {
    promise.push($http.get('/api/skill/' + this.id))
  }

  Promise.all(promise)
    .then((response) => {
      this.object = response[0].data
      this.object.photo = "../../img/test.jpg";
      if (this.type === "friend"){
        this.object.name = this.object.firstname + " " + this.object.lastname;
      }
      if (!$scope.$$phase) {
        $scope.$apply();
      }
    })
  .catch(() => alert('Cannot load informations'))

// this.object = {
//     name: "TEST NAME",
//     description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
//     date: new Date(),
//     type: this.type,
//     photo: "../../img/test.jpg",
// }

this.changeLocation = (route) => {
  console.log("change location")
  $location.path('/' + route)
}

})