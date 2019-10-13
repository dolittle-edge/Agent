/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package log

import (
	original "log"
)

// Debugln prints a debug message using the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Debugln(v ...interface{}) {
	original.Println(append([]interface{}{"[DBG]"}, v...)...)
}

// Debugf prints a debug message using the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	original.Printf("[DBG] "+format, v...)
}

// Informationln prints an information message using the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Informationln(v ...interface{}) {
	original.Println(append([]interface{}{"[INFO]"}, v...)...)
}

// Informationf prints an information message using the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Informationf(format string, v ...interface{}) {
	original.Printf("[INFO] "+format, v...)
}

// Warningln prints an information message using the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Warningln(v ...interface{}) {
	original.Println(append([]interface{}{"[WARN]"}, v...)...)
}

// Warningf prints an information message using the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, v ...interface{}) {
	original.Printf("[WARN] "+format, v...)
}

// Errorln prints an information message using the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Errorln(v ...interface{}) {
	original.Println(append([]interface{}{"[ERR]"}, v...)...)
}

// Errorf prints an information message using the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	original.Printf("[ERR] "+format, v...)
}
