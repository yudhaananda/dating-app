ALTER TABLE user_activities 
ADD FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE user_activities 
ADD FOREIGN KEY (passed_user_id) REFERENCES users(id);

ALTER TABLE user_activities 
ADD FOREIGN KEY (liked_user_id) REFERENCES users(id);

ALTER TABLE users 
ADD FOREIGN KEY (premium_feature_id) REFERENCES premium_features(id);
