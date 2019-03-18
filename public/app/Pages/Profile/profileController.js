'use strict';

var Profile = angular.module('Profile', ['ngRoute'])

Profile.config(['$routeProvider', function ($routeProvider) {
  $routeProvider.when('/Profile', {
    templateUrl: 'Pages/Profile/profile.html',
    controller: 'ProfileController',
    controllerAs: '$ctrl',
  });
}])

/**
 * @brief Profile controller
 */
Profile.controller('ProfileController', function ($location, $http, $scope) {
  console.log("hello dashboar controller")
  this.userId = localStorage.getItem("id");
  document.getElementById('test').style.display = "";
  this.hobbies = [];
  this.companies = [];
  this.skills = [];
  this.friends = [];
  this.jobs = [];


  this.me = undefined;

  $http.get('/api/profile/' + this.userId)
    .then((response) => response.data)
    .then((response) => {
      console.log(response);
      this.me = response

      const promises = []

      promises.push($http.get('/api/profile/' + this.userId + "/companies"))
      promises.push($http.get('/api/profile/' + this.userId + "/hobbies"))
      promises.push($http.get('/api/profile/' + this.userId + "/skills"))
      promises.push($http.get('/api/profile/' + this.userId + "/followed"))
      promises.push($http.get('/api/profile/' + this.userId + "/jobs"))

      Promise.all(promises)
        .then((responses) => {
          responses = responses.map(elem => elem.data)
          // this.companies = responses[0];
          this.hobbies = responses[1];
          // this.skills = responses[2];
          // this.friends = response[3];
          // this.jobs = response[4];
          this.hobbies.forEach((elem) => {
            elem.photo = "../../img/test.jpg"
          })
          if (!$scope.$$phase) {
            $scope.$apply();
          }
          console.log(this.hobbies)
        }).catch((err) => alert("ERROR" + err))
    })



  this.changePasswordBool = false;

  this.clickModify = () => {
    this.changePasswordBool = !this.changePasswordBool
  }

  /**
   * @brief validation changement mot de passe handler button
   */
  this.validateChangePassword = () => {
    this.changePasswordBool = false;
  }

  for (let i = 0; i < 10; i++) {
    this.friends.push({
      name: "toto",
      surname: "titi " + i,
      age: "100",
      city: "Paris",
      photo: "../../img/test.jpg"
    })
  }
  
  for (let i = 0; i < 10; i++) {
    this.companies.push({
      name: "Apple",
      photo: "../../img/test.jpg",
    })
  }

  for (let i = 0; i < 10; i++) {
    this.skills.push({
      name: "C++",
      photo: "../../img/test.jpg",
    })
  }

  for (let i = 0; i < 10; i++) {
    this.jobs.push({
      name: "C++",
      photo: "../../img/test.jpg",
    })
  }

  /**
   * @brief change location page for detail page
   */
  this.moveToDetails = (type) => {
    console.log("move to details")
    $location.path('/Details/' + type)
  }

  /**
  * @brief follow handler button
  */
  this.followClick = (obj) => {
    console.log("follow", obj)
  }

  /**
   * @brief unfollow handler button
   */
  this.unfollowClick = (obj) => {
    console.log("unfollow", obj)
  }
});