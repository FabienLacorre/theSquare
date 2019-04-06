'use strict';

// Declare app level module which depends on views, and core components
angular.module('myApp', [
  'ngRoute',
  'ui.router',
  'navBar',
  'Dashboard',
  'Connection',
  'Configuration',
  'SignIn',
  'Profile',
  'Details',
  'Proposition',
]).
config(['$routeProvider', function($routeProvider) {
  $routeProvider.otherwise({redirectTo: '/'});
}]);
