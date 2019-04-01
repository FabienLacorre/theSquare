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
Dashboard.controller('DashboardController', function ($location, $http) {
  console.log("hello dashboar controller")
  document.getElementById('test').style.display = "";

  console.log(localStorage.getItem("id"));
  this.dashboardObjects = []
  this.pattern = ""
  this.searchValue = "all"

  this.submitSearch = () => {
    /*$http.get(`/api/${this.searchValue}/search/${this.pattern}`).then((response) => response.data)
    .then((response) => {
      console.log(this.searchValue)
      console.log(response)
      this.dashboardObjects = response
      this.dashboardObjects = this.dashboardObjects.map((elem) => {
        return {
          nom: elem.name,
          photo: "../../img/devlogo.png",
          description: elem.description,
        }
      })
    }).catch(() => alert("ERROR REQUEST"))*/

    this.createdashboad(`/api/${this.searchValue}/search/${this.pattern}`)
  }

  // TODO replace copany by all

  this.createdashboad = (route) => {
    $http.get(route).then((response) => response.data)
    .then((response) => {
      console.log(this.searchValue)
      console.log(response)
      this.dashboardObjects = response
      this.finalObjects = []
      this.dashboardObjects.companies = this.dashboardObjects.companies.map((elem) => {
        return {
          nom: elem.name,
          photo: "../../img/devlogo.png",
          description: elem.description,
          type: 'companie',
          id: elem.id
        }
      })
      this.dashboardObjects.companies.forEach(element => {this.finalObjects.push(element)})    
  
      this.dashboardObjects.jobs = this.dashboardObjects.jobs.map((elem) => {
        return {
          nom: elem.name,
          photo: "../../img/devlogo.png",
          description: elem.description,
          type: 'job',
          id: elem.id
        }
      })
      this.dashboardObjects.jobs.forEach(element => {this.finalObjects.push(element)})    
  
      this.dashboardObjects.profiles = this.dashboardObjects.profiles.map((elem) => {
        return {
          nom: elem.firstname + " " + elem.lastname,
          photo: "../../img/devlogo.png",
          description: elem.description,
          type: 'friend',
          id: elem.id
        }
      })
      this.dashboardObjects.profiles.forEach(element => {this.finalObjects.push(element)})    
  
      this.dashboardObjects.skills = this.dashboardObjects.skills.map((elem) => {
        return {
          nom: elem.name,
          photo: "../../img/devlogo.png",
          description: elem.description,
          type: 'skill',
          id: elem.id
        }
      })
      this.dashboardObjects.skills.forEach(element => {this.finalObjects.push(element)})    
  
      this.dashboardObjects.hobbies = this.dashboardObjects.hobbies.map((elem) => {
        return {
          nom: elem.name,
          photo: "../../img/devlogo.png",
          description: elem.description,
          type: 'hobbie',
          id: elem.id
        }
      })
      this.dashboardObjects.hobbies.forEach(element => {this.finalObjects.push(element  )})    

      this.dashboardObjects = this.finalObjects
  
      console.log(this.dashboardObjects)
    }).catch(() => alert("ERROR REQUEST"))
  }

  this.createdashboad(`/api/company/search/`);

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