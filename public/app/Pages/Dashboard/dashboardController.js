'use strict';

var Dashboard = angular.module('Dashboard', ['ngRoute'])

Dashboard.config(['$routeProvider', function ($routeProvider) {
  $routeProvider.when('/Dashboard', {
    templateUrl: 'Pages/Dashboard/dashboard.html',
    controller: 'DashboardController',
    controllerAs: '$ctrl',
  });
}])

/**
 * @brief Dashboard controller
 */
Dashboard.controller('DashboardController', function ($location) {
  console.log("hello dashboar controller")
  document.getElementById('test').style.display = "";

  console.log(localStorage.getItem("id"));
  this.dashboardObjects = []

  /**
   * Temporaire avant d'avoir les get de la base
   */
  for (let i = 0; i < 4; i++) {
    this.dashboardObjects.push({
      nom: "APPLE",
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
      date: new Date(),
      type: "companie",
      photo: "../../img/applelogo.png",
    });
    this.dashboardObjects.push({
      nom: "FULL STACK DEV",
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
      date: new Date(),
      type: "job",
      photo: "../../img/devlogo.png",
    });
    this.dashboardObjects.push({
      nom: "C / C++",
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
      date: new Date(),
      type: "hobbie",
      photo: "../../img/clogo.png",
    });
    this.dashboardObjects.push({
      nom: "TOTO",
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
      date: new Date(),
      type: "friend",
      photo: "../../img/clogo.png",
    });
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

  /**
   * @brief change location page for detail page
   */
  this.moveToDetails = (type, id) => {
    console.log("move to details " + type)
    $location.path('/Details/' + type + "/" + id)
  }
});