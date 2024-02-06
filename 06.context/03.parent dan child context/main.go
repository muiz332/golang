/*

- parent dan child context


context itu dia menganut yang namanya parent dan child
artinya saat kita membuat context kita bisa membuat child context dari context yang sudah kita buat

parent context itu bisa memiliki banyak child akan tetapi child itu hanya dapat memiliki
satu context saja

konsep ini mirip dengan inherintace pada bahasa pemrograman yang berorientasi object


context A  => context B => context D, context E
context A => context C => context F

misalkan saya memiliki context A yang memiliki dua buah child yaitu context B dan context C
dan context B itu memiliki 2 buah child lagi yaitu context D dan context E

dan context C memiliki satu child yaitu context F

berarti context B dan C itu parentnya adalah context A
context D dan context E itu parentnya adalah context B dan juga context A

jadi saling terhubung, nah ketika kalian bikin childnya, semua fitur yang dimiliki dari context parent
itu akan diwariskan ke semua childnya

jadi misalkan nanti context A memiliki fitur cancel, timeout dan ketika kalian bikin child context B
maka fitur itu akan diwariskan ke context B serta child childnya

misalkan nanti context A itu melakukan pembatalan, maka semua childnya juga akan ikut dibatalkan
kalo misalkan saya mengcancel yang context B saja maka context D dan context E akan
ikut dibatalkan

jadi parent dari context B tidak akan ikut dibatalkan ya
dan itu berlaku juga kalo kita menambahkan data pada contextnya



- immutable

context merupakan object yang immutable, artinya setelah context dibuat, dia tidak bisa diubah lagi
oleh karena itu context itu aman ketika dikirim ke beberapa proses

nah kalo kalian coba untuk menambahkan value kedalam context yang sudah ada
maka secara otomatis akan membentuka child context baru, bukan merubah context tersebut





*/

package main

func main() {

}
