module.exports = function (grunt) {
	grunt.initConfig({
		nodewebkit: {
			options: {
				version: '0.10.5',
				buildDir: './bin',
				platforms: ['linux64']
			},
			src: './app/**/*'
		},
		compass: {
			dist: {
				options: {
					basePath: 'app/styles',
					config: 'app/styles/config.rb'
				}
			}
		},
		jshint: {
			all: ['Gruntfile.js', 'app/*.js']
		},
		watch: {
			scripts: {
				files: ['app/styles/sass/*'],
				tasks: ['compass'],
			},
		},
	});

	grunt.loadNpmTasks('grunt-node-webkit-builder');
	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-compass');
	grunt.loadNpmTasks('grunt-contrib-watch');

	grunt.registerTask('build', ['compass', 'nodewebkit']);
};
