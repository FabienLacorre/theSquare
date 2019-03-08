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
Profile.controller('ProfileController', function ($location) {
  console.log("hello dashboar controller")
  document.getElementById('test').style.display = "";

  this.changePasswordBool = false;

  this.clickModify = () => {
    this.changePasswordBool = !this.changePasswordBool
  }

  /**
   * @brief validation changement mot de passe handler button
   */
  this.validateChangePassword = () => {
    console.log("MODIFICATION PASSWORD CLICK")
    this.changePasswordBool = false;
  }

  this.me = {
    name: "LACORRE",
    surname: "Fabien",
    age: 25,
    city: "Rennes",
    friends: [],
  }

  this.friends = []
  for (let i = 0; i < 10; i++) {
    this.friends.push({
      name: "toto",
      surname: "titi " + i,
      age: "100",
      city: "Paris",
      photo: "../../img/test.jpg"
    })
  }

  this.hobbies = []
  for (let i = 0; i < 10; i++) {
    this.hobbies.push({
      name: "hobbies test " + i,
      photo: "../../img/test.jpg"
    })
  }

  this.companies = []
  for (let i = 0; i < 10; i++) {
    this.companies.push({
      name: "Apple",
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