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
          console.log(responses)
          this.companies = responses[0];
          this.companies.forEach((elem) => {
            elem.photo = "../../img/test.jpg"
          })
          this.hobbies = responses[1];
          this.hobbies.forEach((elem) => {
            elem.photo = "../../img/test.jpg"
          })
          this.skills = responses[2];
          this.skills.forEach((elem) => {
            elem.photo = "../../img/test.jpg"
          })
          this.friends = responses[3];
          this.friends.forEach((elem) => {
            elem.photo = "../../img/test.jpg"
          })
          this.jobs = responses[4];
          this.hobbies.forEach((elem) => {
            elem.photo = "../../img/test.jpg"
          })
          if (!$scope.$$phase) {
            $scope.$apply();
          }
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

  /**
   * @brief change location page for detail page
   */
  this.moveToDetails = (type, id) => {
    console.log("move to details")
    $location.path('/Details/' + type + "/" + id)
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