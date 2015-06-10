var request = require('request');

App = {
	Models: {},
	Collections: {},
	Views: {}
};

template = function(id) {
	return _.template($('#' + id + '-template').html());
};

// overriding backbone sync to use request library
Backbone.sync = function(method, model, options) {
	var type = methodMap[method];
	// Default options, unless specified.
	_.defaults(options || (options = {}), {
		emulateHTTP: Backbone.emulateHTTP,
		emulateJSON: Backbone.emulateJSON
	});
	// Default JSON-request options.
	var params = {type: type, dataType: 'json', headers: {}};
	// Ensure that we have a URL.
	if (!options.url) {
		params.url = _.result(model, 'url') || urlError();
	}
	// Ensure that we have the appropriate request data.
	if (options.data === null && model && (method === 'create' || method === 'update' || method === 'patch')) {
		params.contentType = 'application/json';
		params.data = JSON.stringify(options.attrs || model.toJSON(options));
	}
	// For older servers, emulate JSON by encoding the request into an HTML-form.
	if (options.emulateJSON) {
		params.contentType = 'application/x-www-form-urlencoded';
		params.data = params.data ? {model: params.data} : {};
	}
	// For older servers, emulate HTTP by mimicking the HTTP method with `_method`
	// And an `X-HTTP-Method-Override` header.
	if (options.emulateHTTP && (type === 'PUT' || type === 'DELETE' || type === 'PATCH')) {
		params.type = 'POST';
		if (options.emulateJSON) params.data._method = type;
		params.headers['X-HTTP-Method-Override'] = type;
	}
	// Don't process data on a non-GET request.
	if (params.type !== 'GET' && !options.emulateJSON) {
		params.processData = false;
	}

	request({ url: params.url, json: true, method: params.type, headers: params.headers }, function (err, result, body) {
		if (err) {
			return options.error(err);
		}
		return options.success(body);
	});
};

// Map from CRUD to HTTP for our default `Backbone.sync` implementation.
var methodMap = {
	'create': 'POST',
	'update': 'PUT',
	'patch': 'PATCH',
	'delete': 'DELETE',
	'read': 'GET'
};

// dev
var gulp = require('gulp');
gulp.task('reload', function () {
	if (location) location.reload();
});

gulp.watch('**/*', ['reload']);
