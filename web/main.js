class DigitalClock {
    constructor(element) {
        this.element = element;
    }

    getTimeParts() {
        const now = new Date();

        return {
            hour: now.getHours() % 12 || 12
            minute: now.getMinutes(), 
            isAM: now.getHours() < 12
        };
    }
}

const clockElement = document.querySelector(".clock");
const clockObject = new DigitalClock(clockElement);