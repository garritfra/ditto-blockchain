const gulp = require("gulp");
const ts = require("gulp-typescript");

gulp.task("typescript", function() {
  return gulp
    .src("./src/*.ts")
    .pipe(ts())
    .pipe(gulp.dest("dist"));
});

gulp.task("build", ["typescript"]);
gulp.task("default", ["build"]);
