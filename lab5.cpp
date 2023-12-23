#include <iostream>
#include <vector>
#include <memory>
#include <algorithm>
#include <unordered_set>
#include <stdexcept>

// Абстрактный класс персонажа
class Character {
public:
	virtual ~Character() {}
	virtual void setName(const std::string& name) = 0;
	virtual void setElement(const std::string& element) = 0;
	virtual void setRole(const std::string& role) = 0;
	virtual void setAttributes(const std::string& attributes) = 0;
	virtual void displayInfo() = 0;
	virtual std::string getName() = 0;
protected:
	std::string name;
	std::string element;
	std::string role;
};

// Абстрактная фабрика персонажей
class CharacterFactory {
public:
	virtual Character* createMage() = 0;
	virtual Character* createWarrior() = 0;
};

// Конкретный класс персонажа - Маг
class Mage : public Character {
private:
	std::string attributes;

public:
	void setName(const std::string& name) override {
		this->name = name;
	}

	void setElement(const std::string& element) override {
		this->element = element;
	}

	void setRole(const std::string& role) override {
		this->role = role;
	}

	void setAttributes(const std::string& attributes) override {
		this->attributes = attributes;
	}

	void displayInfo() override {
		std::cout << "Character Info:\n";
		std::cout << "Name: " << name << "\n";
		std::cout << "Element: " << element << "\n";
		std::cout << "Role: " << role << "\n";
		std::cout << "Attributes: " << attributes << "\n";
		std::cout << "----------------------" << attributes << "\n";
		std::cout << "" << "\n";
	}

	std::string getName() override {
		return name;
	}
};

// Конкретный класс персонажа - Воин
class Warrior : public Character {
private:
	std::string attributes;

public:
	void setName(const std::string& name) override {
		this->name = name;
	}

	void setElement(const std::string& element) override {
		this->element = element;
	}

	void setRole(const std::string& role) override {
		this->role = role;
	}

	void setAttributes(const std::string& attributes) override {
		this->attributes = attributes;
	}

	void displayInfo() override {
		std::cout << "Character Info:\n";
		std::cout << "Name: " << name << "\n";
		std::cout << "Element: " << element << "\n";
		std::cout << "Role: " << role << "\n";
		std::cout << "Attributes: " << attributes << "\n";
		std::cout << "----------------------" << attributes << "\n";
		std::cout << "" << "\n";
	}

	std::string getName() override {
		return name;
	}
};

// Конкретная фабрика создания персонажей стихии Земли
class EarthCharFactory : public CharacterFactory {
public:
	Character* createMage() override {
		return new Mage();
	}

	Character* createWarrior() override {
		return new Warrior();
	}
};

// Конкретная фабрика создания персонажей стихии Воды
class WaterCharFactory : public CharacterFactory {
public:
	Character* createMage() override {
		return new Mage();
	}

	Character* createWarrior() override {
		return new Warrior();
	}
};

class CharacterManager {
private:
	std::vector<Character*> characters;
	std::unordered_set<std::string> characterNames;

public:
	bool isNameUnique(const std::string& name) {
		return characterNames.find(name) == characterNames.end();
	}

	void addCharacter(Character* character) {
		if (isNameUnique(character->getName())) {
			characters.push_back(character);
			characterNames.insert(character->getName());
		} else {
			throw std::invalid_argument("Character name must be unique!");
		}
	}
};

int main() {
	CharacterManager manager;

	while (true){
		std::cout << "Какой стихии персонажа вы хотите создать? (1 - Земли, 2 - Воды, 0 - Выход)\n";
		int a;
		std::cin >> a;

		if (a == 1){
		CharacterFactory* earthFactory = new EarthCharFactory();
		std::cout << "Какого класса персонажа вы хотите создать? (1 - Маг, 2 - Воин)\n";
		int b;
		std::cin >> b;
		if (b == 1){
			Character* mage = earthFactory->createMage();
			std::cout << "Введите имя персонажа: \n";
			std::string nm;
			std::cin >> nm;
			std::cout << "Свойства персонажа: \n";
			std::string sv;
			std::cin >> sv;
			mage->setName(nm);
			mage->setElement("Earth");
			mage->setRole("Mage");
			mage->setAttributes(sv);
			std::cout << "";
			mage->displayInfo();
			try {
				manager.addCharacter(mage);
			} catch (const std::exception& e) {
				// обработка ошибки
				std::cout << "Имя должно быть уникальным" << e.what() << std::endl;
			}
		} else if (b == 2){
			Character* warrior = earthFactory->createWarrior();
			std::cout << "Введите имя персонажа: \n";
			std::string nm;
			std::cin >> nm;
			std::cout << "Свойства персонажа: \n";
			std::string sv;
			std::cin >> sv;
			warrior->setName(nm);
			warrior->setElement("Earth");
			warrior->setRole("Warrior");
			warrior->setAttributes(sv);
			std::cout << "";
			warrior->displayInfo();
			try {
				manager.addCharacter(warrior);
			} catch (const std::exception& e) {
				// обработка ошибки
				std::cout << "Имя должно быть уникальным" << e.what() << std::endl;
			}
		}
	} else if (a == 2){
		CharacterFactory* waterFactory = new WaterCharFactory();
		std::cout << "Какого класса персонажа вы хотите создать? (1 - Маг, 2 - Воин)\n";
		int b;
		std::cin >> b;
		if (b == 1){
			Character* mage = waterFactory->createMage();
			std::cout << "Введите имя персонажа: \n";
			std::string nm;
			std::cin >> nm;
			std::cout << "Свойства персонажа: \n";
			std::string sv;
			std::cin >> sv;
			mage->setName(nm);
			mage->setElement("Water");
			mage->setRole("Mage");
			mage->setAttributes(sv);
			std::cout << "";
			mage->displayInfo();
			try {
				manager.addCharacter(mage);
			} catch (const std::exception& e) {
				// обработка ошибки
				std::cout << "Имя должно быть уникальным" << e.what() << std::endl;
			}
		} else if (b == 2){
			Character* warrior = waterFactory->createWarrior();
			std::cout << "Введите имя персонажа: \n";
			std::string nm;
			std::cin >> nm;
			std::cout << "Свойства персонажа: \n";
			std::string sv;
			std::cin >> sv;
			warrior->setName(nm);
			warrior->setElement("Water");
			warrior->setRole("Warrior");
			warrior->setAttributes(sv);
			std::cout << "";
			warrior->displayInfo();
			try {
				manager.addCharacter(warrior);
			} catch (const std::exception& e) {
				// обработка ошибки
				std::cout << "Имя должно быть уникальным" << e.what() << std::endl;
			}

		}
	} else if (a == 0){
		break;
	}
	std::cout << std::endl;
	}
	return 0;
}
