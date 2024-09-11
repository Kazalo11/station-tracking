FROM node:16.17.0-bullseye-slim
WORKDIR /app
COPY /web .
RUN npm install
RUN npm run build
EXPOSE 3000
CMD ["npm", "run", "preview"]