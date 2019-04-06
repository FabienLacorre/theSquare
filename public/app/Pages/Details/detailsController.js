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
    promise.push($http.get('/api/profile/' + localStorage.getItem("id") + "/companies/" + this.id))
  } else if (this.type === "friend") {
    promise.push($http.get('/api/profile/' + this.id))
    promise.push($http.get('/api/profile/' + localStorage.getItem("id") + "/follow/" + this.id))
  } else if (this.type === "job") {
    promise.push($http.get('/api/job/' + this.id))
    promise.push($http.get('/api/profile/' + localStorage.getItem("id") + "/jobs/" + this.id))
  } else if (this.type === "hobbie") {
    promise.push($http.get('/api/hobby/' + this.id))
    promise.push($http.get('/api/profile/' + localStorage.getItem("id") + "/hobbies/" + this.id))
  } else if (this.type === "skill") {
    promise.push($http.get('/api/skill/' + this.id))
    promise.push($http.get('/api/profile/' + localStorage.getItem("id") + "/skills/" + this.id))
  }

  Promise.all(promise)
    .then((response) => {
      this.isLike = response[1].data
      this.object = response[0].data


      this.object.photo = "../../img/test.jpg";
      if (this.type === "friend"){
        this.object.name = this.object.firstname + " " + this.object.lastname;
        this.object.photo = this.object.image != null && this.object.image !== "" ? this.object.image : "../../img/test.jpg";
        
        const promises = []
        promises.push($http.get('/api/profile/' + this.object.id + "/companies"))
        promises.push($http.get('/api/profile/' + this.object.id + "/hobbies"))
        promises.push($http.get('/api/profile/' + this.object.id + "/skills"))
        promises.push($http.get('/api/profile/' + this.object.id + "/followed"))
        promises.push($http.get('/api/profile/' + this.object.id + "/jobs"))

        Promise.all(promises)
        .then((responses) => {
          responses = responses.map(elem => elem.data)
          this.friendCompanies = responses[0]
          this.friendHobbies = responses[1]
          this.friendSkills = responses[2]
          this.friendFollowed = responses[3]
          console.log(this.friendSkills)
          this.friendJobs = responses[4]
          if (!$scope.$$phase) {
            $scope.$apply()
          }
        }).catch(() => alert("Error loading data"))

      }
      if (this.type === "companie"){
        this.object.photo = this.object.image != null && this.object.image !== "" ? this.object.image : "../../img/test.jpg";
        const promises = [
          $http.get("/api/company/" + this.object.id + "/likers"),
          $http.get("/api/company/" + this.object.id + "/jobs")
        ]
        Promise.all(promises)
        .then((responses) => {
          responses = responses.map(elem => elem.data)
          console.log(responses)
          this.companyLikers = responses[0];
          this.companyJobs = responses[1];
          if (!$scope.$$phase) {
            $scope.$apply()
          }
        }).catch(() => alert("Error loading data"))
      }
      if (this.type === "hobbie"){
        $http.get("/api/hobby/" + this.object.id + "/likers")
        .then((response) => response.data)
        .then((response) => {
          console.log(response)
          this.hobbieLikers = response
          if (!$scope.$$phase) {
            $scope.$apply()
          }
        }).catch(() => alert("Error loading data"))
      }
      if (!$scope.$$phase) {
        $scope.$apply()
      }
      
    })
  .catch(() => alert('Cannot load informations'))

this.changeLocation = (route) => {
  console.log("change location")
  $location.path('/' + route)
}
  /**
   * @brief follow handler button
   */
  this.followClick = (obj) => {
    console.log("follow", obj)
    let route = ""
    if (this.type === "companie"){
      route += "/api/profile/" + localStorage.getItem("id") + "/companies/" + obj.id; 
    }
    if (this.type === "job"){
      route += "/api/profile/" + localStorage.getItem("id") + "/jobs/" + obj.id;
    }
    if (this.type === "friend"){
      route += "/api/profile/" + localStorage.getItem("id") + "/follow/" + obj.id;
    }
    if (this.type === "skill"){
      route += "/api/profile/" + localStorage.getItem("id") + "/skills/" + obj.id;
    }
    if (this.type === "hobbie"){
      route += "/api/profile/" + localStorage.getItem("id") + "/hobbies/" + obj.id;
    }
    $http.post(route)
    .then((response) => response.data)
    .then((response) => {
      console.log(response)
      this.isLike = !this.isLike 
    }).catch(() => alert("ERROR REQUEST"))
  }

  /**
   * @brief unfollow handler button
   */
  this.unfollowClick = (obj) => {
    console.log("unfollow", obj)
    let route = ""
    if (this.type === "companie"){
      route += "/api/profile/" + localStorage.getItem("id") + "/companies/" + obj.id; 
    }
    if (this.type === "job"){
      route += "/api/profile/" + localStorage.getItem("id") + "/jobs/" + obj.id;
    }
    if (this.type === "friend"){
      route += "/api/profile/" + localStorage.getItem("id") + "/follow/" + obj.id;
    }
    if (this.type === "skill"){
      route += "/api/profile/" + localStorage.getItem("id") + "/skills/" + obj.id;
    }
    if (this.type === "hobbie"){
      route += "/api/profile/" + localStorage.getItem("id") + "/hobbies/" + obj.id;
    }
    console.log(route)
    $http.delete(route)
    .then((response) => response.data)
    .then((response) => {
      console.log(response)
      this.isLike = !this.isLike 
    }).catch(() => alert("ERROR REQUEST"))
  }

})