# Golang Logging
log file adalah file yang berisikan informasi kejadian dari sebuah system. logging adalah aksi menambah informasi ke dalam log file tersebut.

Beberapa library golang logging
- Logrus https://github.com/sirupsen/logrus
- Zap https://github.com/uber-go/zap
- Zerolog https://github.com/rs/zerolog

## Logrus
- https://docs.google.com/presentation/d/1edVvzU_sOExvlN4lzYWOxF38v5GAJHdbbYUqsderNO0/edit#slide=id.g10daea00445_0_286
- logger.Println => time, level, message

### Level
dalam logging, hal yang paling penting adalah level. penentuan prioritas atau jenis kejadian.

- Trace
- Debug
- Info (Default level)
- Warn
- Error
- Fatal => os.Exit(1)
- Panic => panic()

```
logger.SetLevel(logrus.TraceLevel)
```

### Output
secara default log dikirim ke console.
```
logger.SetOutput(file)
```

### Formatter
sebelum mengirim data ke output, logrus memformat dulu menggunakan object formatter. secara default ada dua.
- Text Formatter (Default)
- JSON Formatter

```
logger.SetFormatter()
```

### Field dan Fields
saat mengirim informasi ke log, kadang kita ingin menyisipkan sesuatu pada log tersebut.
```
logger.WithField() => One
logger.WithFields() => Many
```

### Entry
sebuah struct representasi dari log yang kita kirimkan. setiap log yang dikirim, akan dibuatkan object Entry.
```
logrus.NewEntry(logger)
```

### Hook
sebuah struct yang bisa kita tambahkan ke logger sebagai callback.
```
logger.AddHook(&SampleHook{})
```

### Singleton
tanpa membuat object new secara otomatis logger akan implementasi std => new

