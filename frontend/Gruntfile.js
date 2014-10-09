module.exports = function (grunt) {
	grunt.initConfig({
		nodewebkit: {
			options: {
				version: '0.10.0',
				buildDir: './build',
				platforms: ['linux64']
			},
			src: './*'
		},
	});

	grunt.loadNpmTasks('grunt-node-webkit-builder');
	grunt.registerTask('default', ['nodewebkit']);
}
