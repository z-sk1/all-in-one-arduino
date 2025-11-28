package main

import (
	_ "embed"

	"log"

	"strconv"

	"strings"

	"github.com/getlantern/systray"
	"github.com/ncruces/zenity"
)

//go:embed icon.ico
var iconData []byte

func setupTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Arduino Control")
	systray.SetTooltip("Control your Arduino from here!")

	systray.SetIcon(iconData)

	var (
		blueLedOn               = false
		greenLedOn              = false
		redLedOn                = false
		yellowLedOn             = false
		buzzOn                  = false
		rgbShowOn               = false
		melodyOn                = false
		sirenOn                 = false
		megalovaniaOn           = false
		portalThemeMainOn       = false
		apertureThemeOn         = false
		overworldThemeOn        = false
		undergroundThemeOn      = false
		servoSpinOn             = false
		servoJoystickOn         = false
		lightsJoystickOn        = false
		buzzerJoystickOn        = false
		buzzingPreciselyOn      = false
		ledMatrixOn             = false
		ledMatrixSmileyOn       = false
		ledMatrixRandomOn       = false
		ledMatrixJoystickGridOn = false
		ledMatrixJoystickAimOn  = false
		ledControlWithPot       = false
		ledControlWithUltra     = false
		ultraReadingOn          = false
		ledMatrixUltraOn        = false
		servoUltraOn            = false
		buzzerUltraOn           = false
		clockworkOn             = false
		lcdOn                   = false
		lcdAutoscrollOn         = false
		lcdServoDebugOn         = false
	)

	// menu items
	sLights := systray.AddMenuItem("Lights", "Lights sections")
	mToggleBlueLED := sLights.AddSubMenuItem("Turn on Blue LED", "Turn on the Blue LED")
	mToggleGreenLED := sLights.AddSubMenuItem("Turn on Green LED", "Turn on the Green LED")
	mToggleRedLED := sLights.AddSubMenuItem("Turn on Red LED", "Turn on the Red LED")
	mToggleYellowLED := sLights.AddSubMenuItem("Turn on Yellow LED", "Turn on the Yellow LED")
	mLightShow := sLights.AddSubMenuItem("Turn on RGB Light Show", "Turn on an RGB Light Show using the LEDS")
	mLightControlJoystick := sLights.AddSubMenuItem("Control the light with Joystick", "Control your light using a joystick")
	mControlLightWithPot := sLights.AddSubMenuItem("Control the brightness with Potentiometer", "Turn on all LEDS and control their brightnesses with a Potentiometer")
	mControlLightWithUltra := sLights.AddSubMenuItem("Control the brightness with Ultrasound", "Turn on all LEDS and control their brightnesses with a Ultrasound Sensor")

	sBuzz := systray.AddMenuItem("Buzzers", "Buzzer section")
	mToggleBuzz := sBuzz.AddSubMenuItem("Turn on Buzz", "Turn on the Buzz")
	mToggleMelody := sBuzz.AddSubMenuItem("Play a Melody!", "Play a simple note sequence")
	mToggleSiren := sBuzz.AddSubMenuItem("Play a Siren", "Play a Siren sound!")
	mToggleMegalovania := sBuzz.AddSubMenuItem("Play MEGALOVANIA", "PLAY MEGALOVANIA :3")
	mBuzzPrecise := sBuzz.AddSubMenuItem("Enter a Precise Frequency to Buzz", "Enter an exact frequency for the buzzer")
	mBuzzControlJoystick := sBuzz.AddSubMenuItem("Control Buzzer using Joystick", "Control the Buzzer using a Joystick")
	mBuzzControlUltra := sBuzz.AddSubMenuItem("Control Buzzer using Ultrasound", "Control the Buzzer using an Ultrasound Sensor")

	sPortal := sBuzz.AddSubMenuItem("Portal Themes", "Portal Themes Section for Buzz")
	mTogglePortalThemeMain := sPortal.AddSubMenuItem("Play Main Theme", "Play the Main Theme of Portal 2")
	mToggleApertureTheme := sPortal.AddSubMenuItem("Play Aperture Science Theme", "Play the Theme of Aperture Science found in Portal 2")

	sMario := sBuzz.AddSubMenuItem("Mario Themes", "Mario Themes Sections for Buzz")
	mToggleOverworldTheme := sMario.AddSubMenuItem("Play Super Mario Bros. Overworld Theme", "Play the iconic 1-1 overworld theme in the og Mario")
	mToggleUndergroundTheme := sMario.AddSubMenuItem("Play Super Mario Bros. Underground Theme", "Play the iconic 1-2 underground them in the og Mario")

	sSilksong := sBuzz.AddSubMenuItem("Silksong Themes", "Hollow Knight: Silksong Themes Sections for Buzz")
	mToggleClockworkTheme := sSilksong.AddSubMenuItem("Play Clockwork Dancers Theme", "Play the Theme for the Clockwork Dancers boss fight")

	sServo := systray.AddMenuItem("Servos", "Servos section")
	mRotate90 := sServo.AddSubMenuItem("Rotate 90 Degrees", "Rotate the servo by 90 Degrees")
	mRotateNeg90 := sServo.AddSubMenuItem("Rotate Negative 90 Degrees", "Rotate the servo by Negative 90 Degrees")
	mToggleSpin := sServo.AddSubMenuItem("Spin the Servo", "Keep on Spinning the Servo")
	mRotatePrecise := sServo.AddSubMenuItem("Enter a Precise Angle to Rotate", "Enter an exact degree to spin the servo")
	mControlJoystick := sServo.AddSubMenuItem("Control the Servo with Joystick", "Control your Servo using a Joystick on the X-Axis")
	mControlUltraServo := sServo.AddSubMenuItem("Control the Servo with Ultrasound Sensor", "Move your hand closer to the sensor to move it to 0 deg, and the opposite to move it closer to 180 deg")

	sMatrix := systray.AddMenuItem("LED Matrix", "LED Matrix Section")
	mToggleMatrix := sMatrix.AddSubMenuItem("Turn on LED Matrix", "Start a loop that turns on and off the matrix")
	mToggleMatrixSmiley := sMatrix.AddSubMenuItem("Turn on Smiley Face", "Draw a Smiley Face on the LED Matrix")
	mCountDownMatrix := sMatrix.AddSubMenuItem("Count Down", "Draw a count down from 9 - 0 on the LED Matrix")
	mRandomMatrix := sMatrix.AddSubMenuItem("Turn on Random LED Matrix", "Make every light turn on and off randomly")
	sControlJoystickMatrix := sMatrix.AddSubMenuItem("Control LED Matrix using Joystick", "Create a bar graph using the x value and y value of the joystick to determine which column and row")
	mControlJoystickMatrixGrid := sControlJoystickMatrix.AddSubMenuItem("Use Grid", "Control using a grid format")
	mControlJoystickMatrixAim := sControlJoystickMatrix.AddSubMenuItem("Use Crosshair", "Control using a crosshair format")
	mControlUltraMatrix := sMatrix.AddSubMenuItem("Control LED Matrix using Ultrasound Sensor", "Get closer to the ultrasound to make the square smaller")

	mClearLEDMatrix := sMatrix.AddSubMenuItem("Clear LED Matrix", "Clear the display and reset the matrix")
	mChangeBrightnessMatrix := sMatrix.AddSubMenuItem("Change Brightness", "Change the brightness to a number between 1-15")

	sUltra := systray.AddMenuItem("Ultrasound Sensor", "Ultrasound Sensor Section")
	mReadUltrasoundDist := sUltra.AddSubMenuItem("Read Ultrasound Distance", "Check the distance recorded by the ultrasound in cm")

	sLCD := systray.AddMenuItem("LCD Display", "LCD Display Section")
	mToggleLCD := sLCD.AddSubMenuItem("Turn on LCD Display", "Turn on the LCD Display and its Backlight")
	mLCDChangeBrightness := sLCD.AddSubMenuItem("Change Brightness", "Change the brightness of the LCD's backlight")

	sLCDPrint := sLCD.AddSubMenuItem("Print Text", "Print Text to the LCD Display")
	mLCDPrintStatic := sLCDPrint.AddSubMenuItem("Static", "Print Static Text")
	mLCDPrintMoving := sLCDPrint.AddSubMenuItem("Autoscrolling", "Print Moving Text")

	sLCDDebug := sLCD.AddSubMenuItem("Debug", "debug with lcd sections")
	mLCDDebugServo := sLCDDebug.AddSubMenuItem("Debug Servo", "Debug the Servo by printing the angle for them in degrees")

	mLCDCursor := sLCD.AddSubMenuItem("Go To", "Guide the cursor anywhere you want on the Display using x and y coords")
	mLCDClear := sLCD.AddSubMenuItem("Clear Display", "Clear the LCD Display and reset everything on it")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Stop Controlling Arduino and Exit")

	go func() {
		for {
			select {
			case <-mToggleBlueLED.ClickedCh:
				if blueLedOn {
					if err := Device.Exec("turnBlueLedOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleBlueLED.SetTitle("Turn on Blue LED")
					blueLedOn = false
				} else {
					if err := Device.Exec("turnBlueLedOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleBlueLED.SetTitle("Turn off Blue LED")
					blueLedOn = true
				}

			case <-mToggleGreenLED.ClickedCh:
				if greenLedOn {
					if err := Device.Exec("turnGreenLedOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleGreenLED.SetTitle("Turn on Green LED")
					greenLedOn = false
				} else {
					if err := Device.Exec("turnGreenLedOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleGreenLED.SetTitle("Turn off Green LED")
					greenLedOn = true
				}

			case <-mToggleRedLED.ClickedCh:
				if redLedOn {
					if err := Device.Exec("turnRedLedOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleRedLED.SetTitle("Turn on Red LED")
					redLedOn = false
				} else {
					if err := Device.Exec("turnRedLedOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleRedLED.SetTitle("Turn off Red LED")
					redLedOn = true
				}

			case <-mToggleYellowLED.ClickedCh:
				if yellowLedOn {
					if err := Device.Exec("turnYellowLedOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleYellowLED.SetTitle("Turn on Yellow LED")
					yellowLedOn = false
				} else {
					if err := Device.Exec("turnYellowLedOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleYellowLED.SetTitle("Turn off Yellow LED")
					yellowLedOn = true
				}

			case <-mToggleBuzz.ClickedCh:
				if buzzOn {
					if err := Device.Exec("turnBuzzOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleBuzz.SetTitle("Turn on Buzz")
					buzzOn = false
				} else {
					if err := Device.Exec("turnBuzzOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleBuzz.SetTitle("Turn off Buzz")
					buzzOn = true
				}

			case <-mLightShow.ClickedCh:
				if rgbShowOn {
					if err := Device.Exec("rgbShowOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mLightShow.SetTitle("Turn on RGB Light Show")
					rgbShowOn = false
				} else {
					if err := Device.Exec("rgbShowOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mLightShow.SetTitle("Turn off RGB Light Show")
					rgbShowOn = true
				}

			case <-mToggleMelody.ClickedCh:
				if melodyOn {
					if err := Device.Exec("melodyOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleMelody.SetTitle("Play a Melody!")
					melodyOn = false
				} else {
					if err := Device.Exec("melodyOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleMelody.SetTitle("Stop playing Melody")
					melodyOn = true
				}

			case <-mToggleSiren.ClickedCh:
				if sirenOn {
					if err := Device.Exec("sirenOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleSiren.SetTitle("Play a Siren")
					sirenOn = false
				} else {
					if err := Device.Exec("sirenOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleSiren.SetTitle("Stop playing Siren")
					sirenOn = true
				}

			case <-mToggleMegalovania.ClickedCh:
				if megalovaniaOn {
					if err := Device.Exec("megalovaniaOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleMegalovania.SetTitle("Play MEGALOVANIA")
					megalovaniaOn = false
				} else {
					if err := Device.Exec("megalovaniaOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleMegalovania.SetTitle("Stop Playing MEGALOVANIA")
					megalovaniaOn = true
				}

			case <-mTogglePortalThemeMain.ClickedCh:
				if portalThemeMainOn {
					if err := Device.Exec("portalMainThemeOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mTogglePortalThemeMain.SetTitle("Play Main Theme")
					portalThemeMainOn = false
				} else {
					if err := Device.Exec("portalMainThemeOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mTogglePortalThemeMain.SetTitle("Stop Playing Main Theme")
					portalThemeMainOn = true
				}

			case <-mToggleApertureTheme.ClickedCh:
				if apertureThemeOn {
					if err := Device.Exec("apertureThemeOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleApertureTheme.SetTitle("Play Aperture Theme")
					apertureThemeOn = false
				} else {
					if err := Device.Exec("apertureThemeOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleApertureTheme.SetTitle("Stop Playing Aperture Theme")
					apertureThemeOn = true
				}

			case <-mToggleOverworldTheme.ClickedCh:
				if overworldThemeOn {
					if err := Device.Exec("overworldThemeOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleOverworldTheme.SetTitle("Play Super Mario Bros. Overworld Theme")
					overworldThemeOn = false
				} else {
					if err := Device.Exec("overworldThemeOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleOverworldTheme.SetTitle("Stop playing Super Mario Bros. Overworld Theme")
					overworldThemeOn = true
				}

			case <-mToggleUndergroundTheme.ClickedCh:
				if undergroundThemeOn {
					if err := Device.Exec("undergroundThemeOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleUndergroundTheme.SetTitle("Play Super Mario Bros. Underground Theme")
					undergroundThemeOn = false
				} else {
					if err := Device.Exec("undergroundThemeOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleUndergroundTheme.SetTitle("Stop playing Super Mario Bros. Underground Theme")
					undergroundThemeOn = true
				}

			case <-mRotate90.ClickedCh:
				if err := Device.Exec("rotateServo90"); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mRotateNeg90.ClickedCh:
				if err := Device.Exec("rotateServo-90"); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mToggleSpin.ClickedCh:
				if servoSpinOn {
					if err := Device.Exec("servoSpinOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleSpin.SetTitle("Spin the Servo")
					servoSpinOn = false
				} else {
					if err := Device.Exec("servoSpinOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleSpin.SetTitle("Stop spinning the Servo")
					servoSpinOn = true
				}

			case <-mRotatePrecise.ClickedCh:
				deg, err := zenity.Entry("Enter an angle for the servo to rotate: (0-180)")
				if err != nil && !strings.Contains(err.Error(), "dialog canceled") {
					log.Fatal(err)
				}

				intDeg, err := strconv.Atoi(deg)
				if err != nil {
					log.Println(err)
					return
				}

				if intDeg >= 180 {
					deg = "180"
				}

				if intDeg <= 0 {
					deg = "0"
				}

				deg = deg + "\n"

				if err := Device.Execf("rotatePrecise %s", deg); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mControlJoystick.ClickedCh:
				if servoJoystickOn {
					if err := Device.Exec("servoJoyControlOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlJoystick.SetTitle("Control Servo with Joystick")
					servoJoystickOn = false
				} else {
					if err := Device.Exec("servoJoyControlOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlJoystick.SetTitle("Stop Controlling Servo with Joystick")
					servoJoystickOn = true
				}

			case <-mLightControlJoystick.ClickedCh:
				if lightsJoystickOn {
					if err := Device.Exec("lightJoyControlOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mLightControlJoystick.SetTitle("Control Lights with Joystick")
					lightsJoystickOn = false
				} else {
					if err := Device.Exec("lightJoyControlOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mLightControlJoystick.SetTitle("Stop Controlling Lights with Joystick")
					lightsJoystickOn = true
				}

			case <-mBuzzControlJoystick.ClickedCh:
				if buzzerJoystickOn {
					if err := Device.Exec("buzzerJoyControlOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mBuzzControlJoystick.SetTitle("Control Buzzer with Joystick")
					buzzerJoystickOn = false
				} else {
					if err := Device.Exec("buzzerJoyControlOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mBuzzControlJoystick.SetTitle("Stop Controlling Buzzer with Joystick")
					buzzerJoystickOn = true
				}

			case <-mBuzzPrecise.ClickedCh:
				if buzzingPreciselyOn {
					freq, err := zenity.Entry("Enter a precise frequency to Buzz: (100, 5000)")
					if err != nil && !strings.Contains(err.Error(), "dialog canceled") {
						log.Fatal(err)
					}

					intFreq, err := strconv.Atoi(freq)
					if err != nil {
						log.Println(err)
					}

					if intFreq >= 5000 {
						freq = "5000"
					}

					if intFreq <= 100 {
						freq = "100"
					}

					if err := Device.Execf("buzzPrecise %s", freq); err != nil {
						log.Printf("Failed to send command: %v", err)
					}
					buzzingPreciselyOn = true
				} else {
					if err := Device.Exec("buzzPreciseOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					buzzingPreciselyOn = false
				}

			case <-mToggleMatrix.ClickedCh:
				if ledMatrixOn {
					if err := Device.Exec("ledMatrixOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleMatrix.SetTitle("Turn on LED Matrix")

					ledMatrixOn = false
				} else {
					if err := Device.Exec("ledMatrixOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleMatrix.SetTitle("Turn off LED Matrix")

					ledMatrixOn = true
				}

			case <-mToggleMatrixSmiley.ClickedCh:
				if ledMatrixSmileyOn {
					if err := Device.Exec("ledMatrixSmileyOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleMatrixSmiley.SetTitle("Turn off Smiley Face")

					ledMatrixSmileyOn = false
				} else {
					if err := Device.Exec("ledMatrixSmileyOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleMatrixSmiley.SetTitle("Turn on Smiley Face")

					ledMatrixSmileyOn = true
				}

			case <-mCountDownMatrix.ClickedCh:
				if err := Device.Exec("ledMatrixCountdown"); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mRandomMatrix.ClickedCh:
				if ledMatrixRandomOn {
					if err := Device.Exec("ledMatrixRandomOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mRandomMatrix.SetTitle("Turn on Random LED Matrix")

					ledMatrixRandomOn = false
				} else {
					if err := Device.Exec("ledMatrixRandomOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mRandomMatrix.SetTitle("Turn off Random LED Matrix")

					ledMatrixRandomOn = true
				}

			case <-mControlJoystickMatrixGrid.ClickedCh:
				if ledMatrixJoystickGridOn {
					if err := Device.Exec("ledMatrixJoyControlGridOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlJoystickMatrixGrid.SetTitle("Control LED Matrix using Joystick")

					ledMatrixJoystickGridOn = false
				} else {
					if err := Device.Exec("ledMatrixJoyControlGridOn"); err != nil {
						log.Printf("Faled to send command: %v", err)
					}

					mControlJoystickMatrixGrid.SetTitle("Stop Controlling LED Matrix using Joystick")

					ledMatrixJoystickGridOn = true
				}

			case <-mClearLEDMatrix.ClickedCh:
				if err := Device.Exec("ledMatrixClear"); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mChangeBrightnessMatrix.ClickedCh:
				brightness, err := zenity.Entry("Enter Brightness for the LED Matrix: (1-15)")
				if err != nil && !strings.Contains(err.Error(), "dialog canceled") {
					log.Fatal(err)
				}

				brightnessInt, err := strconv.Atoi(brightness)
				if err != nil {
					log.Println(err)
					return
				}

				if brightnessInt <= 1 {
					brightness = "1"
				}

				if brightnessInt >= 15 {
					brightness = "15"
				}

				if err := Device.Execf("ledMatrixBrightness %s", brightness); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mControlLightWithPot.ClickedCh:
				if ledControlWithPot {
					if err := Device.Exec("ledPotControlOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlLightWithPot.SetTitle("Control the brightness with Potentiometer")

					ledControlWithPot = false
				} else {
					if err := Device.Exec("ledPotControlOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlLightWithPot.SetTitle("Stop Controlling the brightness with Potentiometer")

					ledControlWithPot = true
				}

			case <-mControlJoystickMatrixAim.ClickedCh:
				if ledMatrixJoystickAimOn {
					if err := Device.Exec("ledMatrixJoyControlAimOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlJoystickMatrixAim.SetTitle("Use Crosshair")

					ledMatrixJoystickAimOn = false
				} else {
					if err := Device.Exec("ledMatrixJoyControlAimOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlJoystickMatrixAim.SetTitle("Stop using Crosshair")

					ledMatrixJoystickAimOn = true
				}

			case <-mControlLightWithUltra.ClickedCh:
				if ledControlWithUltra {
					if err := Device.Exec("ledUltrasoundControlOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlLightWithUltra.SetTitle("Control the brightness with Ultrasound")

					ledControlWithUltra = false
				} else {
					if err := Device.Exec("ledUltrasoundControlOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlLightWithUltra.SetTitle("Stop controlling the brightness with Ultrasound")

					ledControlWithUltra = true
				}

			case <-mReadUltrasoundDist.ClickedCh:
				if ultraReadingOn {
					if err := Device.Exec("ultrasoundReadOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mReadUltrasoundDist.SetTitle("Read Ultrasound Distance")

					ultraReadingOn = false
				} else {
					if err := Device.Exec("ultrasoundReadOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mReadUltrasoundDist.SetTitle("Stop Reading Ultrasound Distance")

					ultraReadingOn = true
				}

			case <-mControlUltraServo.ClickedCh:
				if servoUltraOn {
					if err := Device.Exec("servoUltraControlOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlUltraServo.SetTitle("Control Servo with Ultrasound")

					servoUltraOn = false
				} else {
					if err := Device.Exec("servoUltraControlOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlUltraServo.SetTitle("Stop Controlling Servo with Ultrasound")

					servoUltraOn = true
				}

			case <-mControlUltraMatrix.ClickedCh:
				if ledMatrixUltraOn {
					if err := Device.Exec("ledMatrixUltraControlOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlUltraMatrix.SetTitle("Control LED Matrix with Ultrasound")

					ledMatrixUltraOn = false
				} else {
					if err := Device.Exec("ledMatrixUltraControlOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mControlUltraMatrix.SetTitle("Stop Controlling LED Matrix with Ultrasound")

					ledMatrixUltraOn = true
				}

			case <-mBuzzControlUltra.ClickedCh:
				if buzzerUltraOn {
					if err := Device.Exec("buzzerUltraControlOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mBuzzControlUltra.SetTitle("Control Buzzer with Ultrasound")

					buzzerUltraOn = false
				} else {
					if err := Device.Exec("buzzerUltraControlOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mBuzzControlUltra.SetTitle("Stop Controlling Buzzer with Ultrasound")

					buzzerUltraOn = true
				}

			case <-mToggleClockworkTheme.ClickedCh:
				if clockworkOn {
					if err := Device.Exec("clockworkThemeOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleClockworkTheme.SetTitle("Play Clockwork Dancers Theme")

					clockworkOn = false
				} else {
					if err := Device.Exec("clockworkThemeOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleClockworkTheme.SetTitle("Stop playing Clockwork Dancers Theme")

					clockworkOn = true
				}

			case <-mToggleLCD.ClickedCh:
				if lcdOn {
					if err := Device.Exec("lcdOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleLCD.SetTitle("Turn on LCD Display")

					lcdOn = false
				} else {
					if err := Device.Exec("lcdOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mToggleLCD.SetTitle("Turn off LCD Display")

					lcdOn = true
				}

			case <-mLCDChangeBrightness.ClickedCh:
				if !lcdOn {
					zenity.Error("Turn on LCD Display First!")
					return
				}

				brightness, err := zenity.Entry("Brightness for LCD: (1-255)")
				if err != nil && !strings.Contains(err.Error(), "dialog canceled") {
					log.Fatal(err)
				}

				brightnessInt, err := strconv.Atoi(brightness)
				if err != nil {
					log.Println(err)
					return
				}

				if brightnessInt <= 1 {
					brightness = "1"
				}

				if brightnessInt >= 255 {
					brightness = "255"
				}

				if err := Device.Execf("lcdBrightness %s", brightness); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mLCDPrintStatic.ClickedCh:
				if !lcdOn {
					zenity.Error("Turn on LCD Display First!")
					return
				}

				displayTxt, err := zenity.Entry("Enter text to print on LCD:")
				if err != nil && !strings.Contains(err.Error(), "dialog canceled") {
					log.Fatal(err)
				}

				if err := Device.Execf("lcdPrintStatic %s", displayTxt); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mLCDPrintMoving.ClickedCh:
				if !lcdOn {
					zenity.Error("Turn on LCD Display First!")
					return
				}

				if lcdAutoscrollOn {

					if err := Device.Exec("lcdPrintMovingOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mLCDPrintMoving.SetTitle("Autoscroll")

					lcdAutoscrollOn = false
				} else {
					displayTxt, err := zenity.Entry("Enter text to print on LCD:")
					if err != nil && !strings.Contains(err.Error(), "dialog canceled") {
						log.Fatal(err)
					}

					if err := Device.Execf("lcdPrintMoving %s", displayTxt); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mLCDPrintMoving.SetTitle("Stop Autoscroll")

					lcdAutoscrollOn = true
				}

			case <-mLCDCursor.ClickedCh:
				if !lcdOn {
					zenity.Error("Turn on LCD Display First!")
					return
				}

				cursor, err := zenity.Entry("Cursor row and column: (0-15), e.g (15 4)")
				if err != nil && !strings.Contains(err.Error(), "dialog canceled") {
					log.Fatal(err)
				}

				if err := Device.Execf("lcdGoTo %s", cursor); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mLCDClear.ClickedCh:
				if !lcdOn {
					zenity.Error("Turn on LCD Display First!")
					return
				}

				if err := Device.Exec("lcdClear"); err != nil {
					log.Printf("Failed to send command: %v", err)
				}

			case <-mLCDDebugServo.ClickedCh:
				if !lcdOn {
					zenity.Error("Turn on LCD Display First!")
					return
				}

				if lcdServoDebugOn {
					if err := Device.Exec("lcdDebugServoOff"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mLCDDebugServo.SetTitle("Stop Debugging Servo")

					lcdServoDebugOn = false
				} else {
					if err := Device.Exec("lcdDebugServoOn"); err != nil {
						log.Printf("Failed to send command: %v", err)
					}

					mLCDDebugServo.SetTitle("Debug Servo")

					lcdServoDebugOn = true
				}

			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	if Device != nil {
		Device.Close()
	}
}
