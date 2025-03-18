FROM nginx:latest

# Picks up Build Arguments
ARG MACHINE_PORT=8080   # Default IP if not provided
ARG MACHINE_IP=10.1.0.2 # Default IP if not provided

RUN echo "Machine PORT: ${MACHINE_PORT}"
RUN echo "Machine IP: ${MACHINE_IP}"

# Picks up Environment Values
ENV ENV_MACHINE_PORT=${MACHINE_PORT}
ENV ENV_MACHINE_IP=${MACHINE_IP}

RUN echo "Machine PORT: ${ENV_MACHINE_PORT}"
RUN echo "Machine IP: ${ENV_MACHINE_IP}"

COPY nginx/default.conf /etc/nginx/conf.d/default.conf

# Replace placeholder in default.conf
RUN sed -i "s/HOST_PORT/${MACHINE_PORT}/g" /etc/nginx/conf.d/default.conf
RUN sed -i "s/HOST_IP/${MACHINE_IP}/g" /etc/nginx/conf.d/default.conf

# RUN echo /etc/nginx/conf.d/default.conf
