'use strict';

var Configuration = angular.module('Configuration', ['ngRoute'])

Configuration.config(['$routeProvider', function ($routeProvider) {
  $routeProvider.when('/Configuration', {
    templateUrl: 'Pages/Configuration/configuration.html',
    controller: 'ConfigurationController',
    controllerAs: '$ctrl',
  });
}])

/**
 * @brief Configuration controller
 */
Configuration.controller('ConfigurationController', function () {
  document.getElementById('test').style.display = "";
});