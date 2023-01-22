Есть функция, которая что-то там ищет по файлу. Но делает она это не очень быстро. Надо её оптимизировать.

Задание на работу с профайлером pprof.

Цель задания - научиться работать с pprof, находить горячие места в коде, уметь строить профиль потребления cpu и памяти, оптимизировать код с учетом этой информации. Написание самого быстрого решения не является целью задания.

Для генерации графа вам понадобится graphviz. Для пользователей windows не забудьте добавить его в PATH чтобы была доступна команда dot.

Рекомендую внимательно прочитать доп. материалы на русском (внизу) - там ещё много примеров оптимизации и объяснений как работать с профайлером. Фактически там есть вся информация для выполнения этого задания.

Есть с десяток мест где можно оптимизировать.

Для выполнения задания необходимо чтобы один из параметров ( ns/op, B/op, allocs/op ) был быстрее чем в *BenchmarkSolution* ( fast < solution ) и ещё один лучше *BenchmarkSolution* + 20% ( fast < solution * 1.2), например ( fast allocs/op < 10422*1.2=12506 ).

По памяти ( B/op ) и количеству аллокаций ( allocs/op ) можно ориентироваться ровно на результаты *BenchmarkSolution* ниже, по времени ( ns/op ) - нет, зависит от системы.

Результат в fast.go в функцию FastSearch (изначально там то же самое что в SlowSearch).

Пример результатов с которыми будет сравниваться:
```
$ go test -bench . -benchmem

goos: darwin

goarch: arm64

pkg: hw3_bench

BenchmarkSlow-8               60          20242046 ns/op        20229997 B/op     189840 allocs/op

BenchmarkFast-8              967           1207100 ns/op          488840 B/op       6169 allocs/op

PASS

ok      hw3_bench       2.673s
```

Запуск:
* `go test -v` - чтобы проверить что ничего не сломалось
* `go test -bench . -benchmem` - для просмотра производительности

Советы:
* Смотрите где мы аллоцируем память
* Смотрите где мы накапливаем весь результат, хотя нам все значения одновременно не нужны
* Смотрите где происходят преобразования типов, которые можно избежать
* Смотрите не только на графе, но и в pprof в текстовом виде (list FastSearch) - там прямо по исходнику можно увидеть где что
* Задание предполагает использование easyjson. На сервере эта библиотека есть, подключать можно. Но сгенерированный через easyjson код вам надо поместить в файл с вашей функцией
* Можно сделать без easyjson

Примечание:
* easyjson основан на рефлекции и не может работать с пакетом main. Для генерации кода вам необходимо вынести вашу структуру в отдельный пакет, сгенерить там код, потом забрать его в main


Материалы для чтения

Рефлексия и кодогенерация:
* https://blog.golang.org/laws-of-reflection
* https://habrahabr.ru/post/269887/
* https://golang.org/src/go/ast/example_test.go
* https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go
* https://golang.org/pkg/reflect/
* http://blog.burntsushi.net/type-parametric-functions-golang/
* https://habrahabr.ru/post/269887/
* https://medium.com/kokster/go-reflection-creating-objects-from-types-part-i-primitive-types-6119e3737f5d
* https://medium.com/kokster/go-reflection-creating-objects-from-types-part-ii-composite-types-69a0e8134f20

Производительность:

Материалы на русском:
* https://habrahabr.ru/company/badoo/blog/301990/
* https://habrahabr.ru/company/badoo/blog/324682/
* https://habrahabr.ru/company/badoo/blog/332636/
* https://habrahabr.ru/company/mailru/blog/331784/ - статья про то как Почта@Mail.ru держит 3 миллиона вебсокет-соединений

Материалы на английском:
* https://blog.golang.org/profiling-go-programs
* https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner/ - большая статья, посвященная go tool trace
* https://www.rzaluska.com/blog/important-go-interfaces/
* https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/
* https://github.com/golang/go/wiki/Performance - много про то что можно вытащить из pprof-а
* https://about.sourcegraph.com/go/advanced-testing-in-go/
* https://signalfx.com/blog/a-pattern-for-optimizing-go-2/
* https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
* https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast - вообще в блоге Дейва очень много полезной инфы по го

Тесты:
* https://blog.golang.org/cover - расширенная информация о go test -cover

Полезные инструменты:
* https://mholt.github.io/json-to-go - позволяет по json сформировать структуру на go, в которую он может быть распакован
* https://github.com/mailru/easyjson - кодогенератор для json
