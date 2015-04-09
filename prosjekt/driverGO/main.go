package main

import (
    ".fmt"
)

func main() {
    // Initialize hardware
    if (!Elev_init()) {
        println("Unable to initialize elevator hardware!\n");
        return 1;
    }

    println("Press STOP button to stop elevator and exit program.\n");

    Elev_set_motor_direction(DIRN_UP);

    for {
        // Change direction when we reach top/bottom floor
        if (Elev_get_floor_sensor_signal() == N_FLOORS - 1) {
            Elev_set_motor_direction(DIRN_DOWN);
        } else if (Elev_get_floor_sensor_signal() == 0) {
            Elev_set_motor_direction(DIRN_UP);
        }

        // Stop elevator and exit program if the stop button is pressed
        if (Elev_get_stop_signal()) {
            Elev_set_motor_direction(DIRN_STOP);
            break;
        }
    }

    return 0;
}
