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
    this.all = []
    this.userConnected = undefined
    $http.get('/api/profile/'+ this.userId+ '/propositions/users').then((response) => response.data)
    .then((response) => {
      console.log(response)
      this.userConnected = response
      this.userConnected = this.userConnected.map((elem) => {
        return {
          nom: elem.firstname + " " + elem.lastname,
          photo: "../../img/devlogo.png",
          description: elem.description,
          type: 'friend',
          id: elem.id
        }
      })
      console.log(this.userConnected)
      this.all = this.all.concat(this.userConnected)
      console.log("users", response)

      this.hobbiesConnected = undefined
      $http.get('/api/profile/'+ this.userId+ '/propositions/users/hobbies').then((response) => response.data)
      .then((response) => {
        this.hobbiesConnected = response
        this.hobbiesConnected = this.hobbiesConnected.map((elem) => {
          return {
            nom: elem.name,
            photo: "../../img/devlogo.png",
            description: elem.description,
            type: 'hobbie',
            id: elem.id
          }
        })
        this.all = this.all.concat(this.hobbiesConnected)


        this.companiesConnected = undefined
        $http.get('/api/profile/'+ this.userId+ '/propositions/companies').then((response) => response.data)
        .then((response) => {
          this.companiesConnected = response
          this.companiesConnected = this.companiesConnected.map((elem) => {
            return {
              nom: elem.name,
              photo: "../../img/devlogo.png",
              description: elem.description,
              type: 'companie',
              id: elem.id
            }
          })
          this.all = this.all.concat(this.companiesConnected)
          console.log("companies", response)
          console.log(this.all)
        }).catch((err ) => alert("Error loading datas " + err))

        console.log("hobbies", response)
      }).catch((err) => alert("Error loading datas " + err))

    }).catch((err) => alert("Error loading datas " + err))

    this.followClick = (obj) => {
      console.log("follow", obj)
      let route = ""
      if (obj.type === "companie"){
        route += "/api/profile/" + localStorage.getItem("id") + "/companies/" + obj.id; 
      }
      if (obj.type === "job"){
        route += "/api/profile/" + localStorage.getItem("id") + "/jobs/" + obj.id;
      }
      if (obj.type === "friend"){
        route += "/api/profile/" + localStorage.getItem("id") + "/follow/" + obj.id;
      }
      if (obj.type === "skill"){
        route += "/api/profile/" + localStorage.getItem("id") + "/skills/" + obj.id;
      }
      if (obj.type === "hobbie"){
        route += "/api/profile/" + localStorage.getItem("id") + "/hobbies/" + obj.id;
      }
      $http.post(route)
      .then((response) => response.data)
      .then((response) => {
        console.log(response)
      }).catch(() => alert("ERROR REQUEST"))
    }
  
    /**
     * @brief unfollow handler button
     */
    this.unfollowClick = (obj) => {
      console.log("unfollow", obj)
      let route = ""
      if (obj.type === "companie"){
        route += "/api/profile/" + localStorage.getItem("id") + "/companies/" + obj.id; 
      }
      if (obj.type === "job"){
        route += "/api/profile/" + localStorage.getItem("id") + "/jobs/" + obj.id;
      }
      if (obj.type === "friend"){
        route += "/api/profile/" + localStorage.getItem("id") + "/follow/" + obj.id;
      }
      if (obj.type === "skill"){
        route += "/api/profile/" + localStorage.getItem("id") + "/skills/" + obj.id;
      }
      if (obj.type === "hobbie"){
        route += "/api/profile/" + localStorage.getItem("id") + "/hobbies/" + obj.id;
      }
      $http.delete(route)
      .then((response) => response.data)
      .then((response) => {
        console.log(response)
      }).catch(() => alert("ERROR REQUEST"))
    }
  
    /**
     * @brief change location page for detail page
     */
    this.moveToDetails = (type, id) => {
      console.log("move to details " + type)
      $location.path('/Details/' + type + "/" + id)
    }

  }); 


