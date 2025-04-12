var gulp = require("gulp");
var shell = require('gulp-shell');

// Bu, kaynak değişikliğiyle yeni ikili dosya derler
gulp.task("install-binary", shell.task([
 'go install github.com/ibrahimbayburtlu/Building-Restful-Web-Services-With-Go/chapter-1/romanserver'
]));

// Supervisor'ı yeniden başlat
gulp.task("restart-supervisor", shell.task([
 'supervisorctl restart myserver'
]));

// İkinci argüman, restart-supervisor için install-binary'nin bir bağımlılık olduğunu belirtir
gulp.task('run-sequence', gulp.series('install-binary', 'restart-supervisor'));

gulp.task('watch', function() {
 // Tüm değişiklikler için kaynak kodu izle
 gulp.watch("*", gulp.series('run-sequence'));
});

gulp.task('default', gulp.series('watch')); 