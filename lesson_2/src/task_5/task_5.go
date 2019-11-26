package task_5

import "fmt"

/*"""
5.	Вывести на экран коды и символы таблицы ASCII, начиная с символа
под номером 32 и заканчивая 127-м включительно.
Вывод выполнить в табличной форме: по десять пар "код-символ" в каждой строке.
"""
 */

func AckiiTable(){
	var count = 1
	var sep, rowSep string

	for i:=33; i<127; i++{

		rowSep=""
		if count % 10 == 0 {
			rowSep = "\n"
			sep = ""
		}else if count % 10 == 1{
			sep = "\t"
		}
		fmt.Print(string(i),"-", count, sep, rowSep)
		count++
	}
}