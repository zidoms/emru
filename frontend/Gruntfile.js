module.exports = function (grunt) {
	grunt.initConfig({
		jshint: {
			all: ['Gruntfile.js', 'app/*.js']
		},
		compass: {
			dist: {
				options: {
					basePath: 'app/styles',
					config: 'app/styles/config.rb'
				}
			}
		},
		nodewebkit: {
			options: {
				version: 'latest',
				buildDir: './bin',
				platforms: ['linux64']
			},
			src: './app/**/*'
		},
	});

	grunt.loadNpmTasks('grunt-node-webkit-builder');
	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-compass');

	grunt.registerTask('build', ['compass', 'nodewebkit']);
};
