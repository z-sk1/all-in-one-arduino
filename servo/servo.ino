#include <Servo.h>

const int yellowPin = 11;
const int greenPin = 10;

Servo servo;
int servoPin = 9;

int servoPos;
const int joyX = A0;
const int joyY = A1;
const int joyBtn = 5;

unsigned long lastMoveTime = 0;
int servoStep = 6;  // how many degrees to move each step
int servoDir = 1;

const int btnPin = 8;
int btnState = 0;
bool lastStableBtn = HIGH;
bool lastReading = HIGH;
unsigned long lastDebounceTime = 0;
const unsigned long debounceDelay = 50; 

bool servoJoystickActive = true;

void setup() {
  Serial.begin(9600);

  servoPos = 90;

  servo.attach(servoPin);
  
  servo.write(servoPos);

  pinMode(joyX, INPUT);
  pinMode(joyY, INPUT);
  pinMode(joyBtn, INPUT_PULLUP);

  pinMode(btnPin, INPUT_PULLUP);

  pinMode(yellowPin, OUTPUT);
  pinMode(greenPin, OUTPUT);

  digitalWrite(yellowPin, LOW);
  digitalWrite(greenPin, HIGH);
}

void loop() {
  int reading = digitalRead(btnPin);

  // If reading changed, reset debounce timer
  if (reading != lastReading) {
      lastDebounceTime = millis();
  }
  lastReading = reading;

  // Check if stable for 50ms
  if ((millis() - lastDebounceTime) > debounceDelay) {

      // Has the stable state changed?
      if (reading != lastStableBtn) {
          lastStableBtn = reading;

          // Only toggle when button becomes PRESSED (LOW)
          if (reading == LOW) {
              servoJoystickActive = !servoJoystickActive;
          }
      }
  }

  if (servoJoystickActive) {
    digitalWrite(greenPin, LOW);
    digitalWrite(yellowPin, HIGH);
  } else {
    digitalWrite(yellowPin, LOW);
    digitalWrite(greenPin, HIGH);
  }

  if (servoJoystickActive) {
    int button = digitalRead(joyBtn);

    if (button == LOW) {
      servoSpinNonBlocking();
    } else {
      long totalX = 0;

      const int samples = 5;
      const int center = 512;
      const int deadzone = 200;
      for (int i = 0; i < samples; i++) {
        totalX += analogRead(joyX);
      }

      int avgX = totalX / samples;
      if (abs(avgX - center) < deadzone) avgX = center;

      int posX = map(avgX, 0, 1023, 0, 180);

      servo.write(posX);
    }
  } else {
    servoSpinNonBlocking();
  }

}

void servoSpinNonBlocking() {
    unsigned long currentTime = millis();
    if (currentTime - lastMoveTime >= 25) { // move every 50ms
        lastMoveTime = currentTime;

        servoPos += servoStep * servoDir;

        if (servoPos >= 180) servoDir = -1; // reverse
        if (servoPos <= 0)   servoDir = 1;  // reverse

        servo.write(servoPos);
    }
}